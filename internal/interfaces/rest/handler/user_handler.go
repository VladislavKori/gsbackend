package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/vladislavkori/gsbackend/internal/app/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (s *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Чтение body из Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Декодирование JSON в структуру
	type RequestData struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	var data RequestData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Валидация полей
	var validate = validator.New()

	if err = validate.Struct(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// создания пользователя
	token, id, err := s.service.RegisterUser(data.Email, data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа
	response := map[string]string{
		"access_token": token,
		"user_id":      fmt.Sprintf("%d", id),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
