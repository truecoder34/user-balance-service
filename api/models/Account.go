package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Entity
	Account     User      `json:"account"`
	UserID      uuid.UUID `gorm:"not null" json:"user_id"`
	MoneyAmount uint64    `gorm:"not null" json:"money_amount"`
	Comment     string    `gorm:"size:255;" json:"comment"`
}

func (account *Account) Prepare() {
	account.Entity = Entity{}
	account.MoneyAmount = 0
	account.Comment = html.EscapeString(strings.TrimSpace(account.Comment))
}

/*
	Save new wallet
*/
func (account *Account) SaveAccount(db *gorm.DB) (*Account, error) {
	var err error
	err = db.Debug().Model(&Account{}).Create(&account).Error
	if err != nil {
		return &Account{}, err
	}
	if account.ID != uuid.Nil {
		err = db.Debug().Model(&Account{}).Where("id = ?", account.UserID).Take(&account.Account).Error
		if err != nil {
			return &Account{}, err
		}
	}
	return account, nil
}

/*
	Change account balance
*/
func (account *Account) UpdateAccountBalance(db *gorm.DB, newAmount int64) (*Account, error) {
	var err error
	account.MoneyAmount = uint64(newAmount)
	account.Comment = html.EscapeString(strings.TrimSpace(account.Comment))
	err = db.Debug().Model(&User{}).Where("id = ?", account.ID).Updates(Account{
		MoneyAmount: account.MoneyAmount,
		Comment:     account.Comment,
		// UpdatedAt : time.Now(),
	}).Error
	if err != nil {
		return &Account{}, err
	}
	return account, nil
}

/*
	Get all accounts in DB
*/
func (account *Account) FindAllAccounts(db *gorm.DB) (*[]Account, error) {
	acs := []Account{}
	var err error = db.Debug().Model(&Account{}).Limit(100).Find(&acs).Error
	if err != nil {
		return &[]Account{}, err
	}
	return &acs, nil
}

/*
	TODO: Get all accounts in DB by USER UD
*/

/*
	Delete account by ID
*/
func (account *Account) DeleteAccount(db *gorm.DB, id uuid.UUID) (int64, error) {
	db = db.Debug().Model(&Account{}).Where("id = ?", id).Take(&Account{}).Delete(&Account{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("account entity not found in database")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil

}
