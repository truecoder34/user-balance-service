package dtos

import uuid "github.com/satori/go.uuid"

type ActionDTO struct {
	UserID      uuid.UUID `json:"user_id"`
	MoneyAmount uint64    `json:"money_amount"`
	Action      bool      `json:"action"` // true - add, false - remove
}
