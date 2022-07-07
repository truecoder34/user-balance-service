package dtos

import uuid "github.com/satori/go.uuid"

type OutTransferDTO struct {
	UserIDSender           uuid.UUID `json:"user_id_sender"`
	UserIDReceiver         uuid.UUID `json:"user_id_receiver"`
	NewMoneyAmountSender   uint64    `json:"new_money_amount_sender"`
	NewMoneyAmountReceiver uint64    `json:"new_money_amount_receiver"`
}
