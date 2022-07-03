package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Entity
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Surname     string `gorm:"size:255;not null;unique" json:"surname"`
	Nickname    string `gorm:"size:255;not null;unique" json:"nickname"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	PhoneNumber string `gorm:"size:100;not null;unique" json:"phone_number"`
}

func (user *User) Prepare() {
	user.Entity = Entity{}
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	user.Surname = html.EscapeString(strings.TrimSpace(user.Surname))
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.PhoneNumber = html.EscapeString(strings.TrimSpace(user.PhoneNumber))
}

func (user *User) Validate() error {
	if user.Name == "" {
		return errors.New("name data is required;")
	}
	if user.Surname == "" {
		return errors.New("surname data is required;")
	}
	if user.Nickname == "" {
		return errors.New("nickname data is required;")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid Email")
	}
	if user.PhoneNumber == "" {
		return errors.New("phone number data is required;")
	}
	return nil
}

/*
	Save User entity
*/
func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

/*
	Find all Users entity
*/
func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

/*
	Find User By ID
*/
func (user *User) FindUserByID(db *gorm.DB, uid uuid.UUID) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

/*
	Delete User Entity
*/
func (user *User) DeleteAUser(db *gorm.DB, uid uuid.UUID) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

/*
	Update User Entity
*/
func (user *User) UpdateAUser(db *gorm.DB, uid uuid.UUID) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":         user.Name,
			"surname":      user.Surname,
			"nickname":     user.Nickname,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"update_at":    time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}
