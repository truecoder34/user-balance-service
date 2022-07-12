package dtos

import uuid "github.com/satori/go.uuid"

type BalanceDTO struct {
	UserID   uuid.UUID `json:"user_id"`
	Currency string    `json:"currency"`
}
