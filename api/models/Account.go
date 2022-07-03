package models

// import (
// 	"errors"
// 	"html"
// 	"strings"
// 	"time"

// 	"github.com/jinzhu/gorm"
// )

type Account struct {
	Entity
	Account     User   `json:"Account"`
	MoneyAmount uint64 `gorm:"not null" json:"money_amount"`
}

func (account *Account) Prepare() {
	account.Entity = Entity{}
	account.MoneyAmount = 0
}
