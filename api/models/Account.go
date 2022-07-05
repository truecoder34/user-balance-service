package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	dtos "github.com/truecoder34/user-balance-service/api/DTOs"
)

type Account struct {
	Entity
	Account       User      `json:"account"`
	AccountNumber string    `gorm:"size:22;" json:"account_number"`
	UserID        uuid.UUID `gorm:"not null" json:"user_id"`
	MoneyAmount   uint64    `gorm:"not null;" json:"money_amount"`
	Comment       string    `gorm:"size:255;" json:"comment"`
}

func (account *Account) Prepare() {
	account.Entity = Entity{}
	account.MoneyAmount = 0
	account.AccountNumber = "" // TODO GENERATE 16 or 20 numbers account number
	account.Comment = html.EscapeString(strings.TrimSpace(account.Comment))
	account.Account = User{}
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
func (account *Account) UpdateAccountBalance(db *gorm.DB, action dtos.ActionDTO) (*Account, error) {
	var err error
	// ac := Account{}
	// err = db.Debug().Model(Account{}).Where("user_id = ?", action.UserID).Find(&ac).Error
	// if err != nil {
	// 	return &Account{}, err
	// }

	// if action.Action {

	// }
	// else {

	// }

	//account.MoneyAmount = html.EscapeString(strings.TrimSpace(account.MoneyAmount))
	account.Comment = html.EscapeString(strings.TrimSpace(account.Comment))
	account.AccountNumber = html.EscapeString(strings.TrimSpace(account.AccountNumber))

	err = db.Debug().Model(&User{}).Where("user_id = ?", action.UserID).Updates(Account{
		MoneyAmount:   account.MoneyAmount + action.MoneyAmount,
		Comment:       account.Comment,
		AccountNumber: account.AccountNumber,
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
	if len(acs) > 0 {
		for i, _ := range acs {
			err := db.Debug().Model(&User{}).Where("id = ?", acs[i].UserID).Take(&acs[i].Account).Error
			if err != nil {
				return &[]Account{}, err
			}
		}
	}

	return &acs, nil
}

/*
	Get accounts in DB by USER UD
*/
func (account *Account) FindAccountByUserID(db *gorm.DB, uid uuid.UUID) (*Account, error) {
	var err error
	ac := Account{}
	err = db.Debug().Model(Account{}).Where("user_id = ?", uid).Find(&ac).Error
	if err != nil {
		return &Account{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Account{}, errors.New("Account was not Found")
	}
	return &ac, nil
}

/*
	Find Account By ID
*/
func (account *Account) FindAccountByID(db *gorm.DB, aid uuid.UUID) (*Account, error) {
	var err error
	err = db.Debug().Model(Account{}).Where("id = ?", aid).Take(&account).Error
	if err != nil {
		return &Account{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Account{}, errors.New("Account was not Found")
	}
	return account, err
}

/*
	Get account in DB by ACCOUNT NUMBER
*/
func (account *Account) FindAccountByAccountNumber(db *gorm.DB, accountNumber string) (*Account, error) {
	var err error
	err = db.Debug().Model(Account{}).Where("account_number = ?", accountNumber).Take(&account).Error
	if err != nil {
		return &Account{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Account{}, errors.New("Account was not Found")
	}
	return account, err
}

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
