package dtos

import uuid "github.com/satori/go.uuid"

type TransferDTO struct {
	UserIDSender   uuid.UUID `json:"user_id_sender"`
	UserIDReceiver uuid.UUID `json:"user_id_receiver"`
	MoneyAmount    uint64    `json:"money_amount"`
}
