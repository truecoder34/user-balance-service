package models

import (
	"html"
	"strings"
	"time"

	//"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	dtos "github.com/truecoder34/user-balance-service/api/DTOs"
	"gorm.io/gorm"
)

type Account struct {
	Entity

	AccountNumber string    `gorm:"size:22;" json:"account_number"`
	UserID        uuid.UUID `gorm:"type:uuid;column:user_id;not null" json:"user_id"`
	Account       User      `gorm:"foreignKey:UserID" json:"account"`
	MoneyAmount   uint64    `gorm:"not null;" json:"money_amount"`
	Comment       string    `gorm:"size:255;" json:"comment"`
}

func (account *Account) Prepare() {
	account.Entity = Entity{}
	account.MoneyAmount = 0
	account.AccountNumber = "" // TODO: GENERATE 16 or 20 numbers account number
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
	tx := db.Debug().Model(&Account{}).Session(&gorm.Session{})
	//tx_user := db.Debug().Model(&User{}).Session(&gorm.Session{})
	var actionPattern string
	if action.Action {
		actionPattern = "money_amount + ?"
	} else {
		actionPattern = "money_amount - ?"
	}
	tx.Where("user_id = ?", action.UserID).Take(&Account{}).UpdateColumns(
		map[string]interface{}{
			"money_amount": gorm.Expr(actionPattern, action.MoneyAmount),
			"updated_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Account{}, db.Error
	}
	// This is the display the updated user
	if err = tx.Where("user_id = ?", action.UserID).Take(account).Error; err != nil {
		return &Account{}, err
	}
	if account != nil {
		if err := db.Debug().Model(&User{}).Where("id = ?", action.UserID).Take(&account.Account).Error; err != nil {
			return &Account{}, err
		}
	}

	return account, nil
}

/*
	Transfer money from one account to another
*/
func (account *Account) TransferMoneyBetweenAccounts(db *gorm.DB, transfer dtos.TransferDTO) (Account, Account, error) {
	var err error
	tx := db.Debug().Model(&Account{}).Session(&gorm.Session{})
	// ADD MONEY TO RECEIVER
	tx.Where("user_id = ?", transfer.UserIDReceiver).Take(&Account{}).UpdateColumns(
		map[string]interface{}{
			"money_amount": gorm.Expr("money_amount + ?", transfer.MoneyAmount),
			"updated_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return Account{}, Account{}, db.Error
	}
	// REMOVE MONET FROM SENDER
	tx.Where("user_id = ?", transfer.UserIDSender).Take(&Account{}).UpdateColumns(
		map[string]interface{}{
			"money_amount": gorm.Expr("money_amount - ?", transfer.MoneyAmount),
			"updated_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return Account{}, Account{}, db.Error
	}

	var accountReceiver Account
	var accountSender Account

	if err = tx.Where("user_id = ?", transfer.UserIDReceiver).Take(&accountReceiver).Error; err != nil {
		return Account{}, Account{}, db.Error
	}

	if err = tx.Where("user_id = ?", transfer.UserIDSender).Take(&accountSender).Error; err != nil {
		return Account{}, Account{}, db.Error
	}

	return accountReceiver, accountSender, nil
}

/*
	Get balance of user  by it ID
*/
func (account *Account) GetAccountBalance(db *gorm.DB, uid uuid.UUID) (*Account, error) {
	var err error
	ac := Account{}
	err = db.Debug().Model(Account{}).Where("user_id = ?", uid).Find(&ac).Error
	if err != nil {
		return &Account{}, err
	}

	return &ac, nil
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
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &Account{}, errors.New("Account was not Found")
	// }
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
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &Account{}, errors.New("Account was not Found")
	// }
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
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &Account{}, errors.New("Account was not Found")
	// }
	return account, err
}

/*
	Delete account by ID
*/
func (account *Account) DeleteAccount(db *gorm.DB, id uuid.UUID) (int64, error) {
	db = db.Debug().Model(&Account{}).Where("id = ?", id).Take(&Account{}).Delete(&Account{})

	if db.Error != nil {
		// if gorm.IsRecordNotFoundError(db.Error) {
		// 	return 0, errors.New("account entity not found in database")
		// }
		return 0, db.Error
	}
	return db.RowsAffected, nil

}
