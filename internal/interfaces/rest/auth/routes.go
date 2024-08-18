package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vladislavkori/gsbackend/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type RequestData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type Claims struct {
	email      string    `json:"email"`
	created_at time.Time `json:"created_at"`
	jwt.RegisteredClaims
}

var validate = validator.New()
var jwtSecret = []byte("your_secret_key")

func Router(userRep repository.UserRepository) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		// Чтение body из Request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Декодирование JSON в структуру
		var data RequestData
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Валидация полей
		if err = validate.Struct(data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userRep.FindUserByEmail(data.Email)
		if err != nil {
			http.Error(w, "Error finding user by email", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		if user != nil {
			http.Error(w, "Error, User with this email already exists", http.StatusConflict)
			return
		}

		// Устанавливаем время истечения токена
		expirationTime := time.Now().Add(15 * time.Minute)

		// Создаем клеймы
		claims := &Claims{
			email:      data.Email,
			created_at: time.Now(),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				Issuer:    "your_app_name",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Подписываем токен с использованием секретного ключа
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Error creating JWT", http.StatusConflict)
			fmt.Println("Error creating JWT:", err)
			return
		}

		// Шифруем пароль
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 4)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Создаём пользователя
		id, err := userRep.CreateUser(data.Email, string(hashPassword))
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"access_token": tokenString,
			"user_id":      fmt.Sprintf("%d", id),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's login route, but now it's not implemented"))
	})

	return r
}
