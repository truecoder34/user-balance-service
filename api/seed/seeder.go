package seed

import (
	"github.com/truecoder34/user-balance-service/api/models"
	"gorm.io/gorm"
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
	//var err error
	//	err := db.Debug().DropTableIfExists(&models.Account{}, &models.User{}).Error
	if db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Account{}) {
		db.Debug().Migrator().DropTable("users", "accounts")
		// if err != "" {
		// 	log.Fatalf("cannot drop table: %v", err)
		// }
		db.Debug().AutoMigrate(&models.User{}, &models.Account{})
		// if err != nil {
		// 	log.Fatalf("cannot migrate table: %v", err)
		// }
	}

	for i, _ := range users {
		db.Debug().Model(&models.User{}).Create(&users[i])
		// if err != nil {
		// 	log.Fatalf("cannot seed users table: %v", err)
		// }
		accounts[i].UserID = users[i].ID

		db.Debug().Model(&models.Account{}).Create(&accounts[i])
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}
}
