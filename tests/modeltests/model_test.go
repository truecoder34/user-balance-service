package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/truecoder34/user-balance-service/api/controllers"
	"github.com/truecoder34/user-balance-service/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var server = controllers.Server{}
var userInstance = models.User{}
var accountInstance = models.Account{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		//server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	server.DB.Migrator().DropTable(&models.User{})
	server.DB.AutoMigrate(&models.User{})
	log.Printf("Successfully refreshed table")
	return nil
}

func refreshUserAndAccountTable() error {
	server.DB.Migrator().DropTable(&models.User{})
	server.DB.Migrator().DropTable(&models.Account{})
	server.DB.AutoMigrate(&models.User{}, &models.Account{})
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		Name:        "OneUser",
		Surname:     "OneUser",
		Nickname:    "OneUser",
		Email:       "OneUser@gmail.com",
		PhoneNumber: "+11111111111",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
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
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedOneUserAndOneAccount() (models.Account, error) {
	err := refreshUserAndAccountTable()
	if err != nil {
		return models.Account{}, err
	}
	user := models.User{
		Name:        "OneUser",
		Surname:     "OneUser",
		Nickname:    "OneUser",
		Email:       "OneUser@gmail.com",
		PhoneNumber: "+11111111111",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Account{}, err
	}
	account := models.Account{
		AccountNumber: "1111 2222 3333 4444",
		MoneyAmount:   74770,
		Comment:       "Test 0",
	}
	err = server.DB.Model(&models.Account{}).Create(&account).Error
	if err != nil {
		return models.Account{}, err
	}
	return account, nil
}

func seedUsersAndAccounts() ([]models.User, []models.Account, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Account{}, err
	}
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

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		accounts[i].UserID = users[i].ID

		err = server.DB.Model(&models.Account{}).Create(&accounts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, accounts, nil
}
