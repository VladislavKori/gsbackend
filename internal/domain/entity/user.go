package entity

import "time"

type User struct {
	ID                      int64
	Email                   string
	Password                string
	AvatarURL               *string
	CurrentDeliveryAdressId *int64
	CreatedAt               time.Time
}
