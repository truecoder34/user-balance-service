package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/truecoder34/user-balance-service/api/models"
)

var users = []models.User{
	models.User{
		Name:        "Tom",
		Surname:     "Hardy",
		Nickname:    "GodTom777",
		Email:       "th777@gmail.com",
		PhoneNumber: "+13477778888",
	},
	models.User{
		Name:        "Александр",
		Surname:     "Петров",
		Nickname:    "AlexKing888",
		Email:       "alex.ptrv@yandex.ru",
		PhoneNumber: "+799913831313",
	},
	models.User{
		Name:        "John",
		Surname:     "Doe",
		Nickname:    "JD",
		Email:       "jd@yahoo.com",
		PhoneNumber: "+11111111111",
	},
}

var accounts = []models.Account{
	models.Account{
		AccountNumber: "1111 2222 3333 4444",
		MoneyAmount:   74770,
		Comment:       "Test 0",
	},
	models.Account{
		AccountNumber: "2222 3333 4444 5555",
		MoneyAmount:   74770,
		Comment:       "Test 1",
	},
	models.Account{
		AccountNumber: "3333 4444 5555 6666",
		MoneyAmount:   74770,
		Comment:       "Test 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Account{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Account{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Account{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		accounts[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Account{}).Create(&accounts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
