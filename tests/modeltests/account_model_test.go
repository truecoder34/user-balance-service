package modeltests

import (
	"log"
	"testing"

	"github.com/truecoder34/user-balance-service/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllPosts(t *testing.T) {

	err := refreshUserAndAccountTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}
	_, _, err = seedUsersAndAccounts()
	if err != nil {
		log.Fatalf("Error seeding user and post  table %v\n", err)
	}
	posts, err := accountInstance.FindAllAccounts(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 3)
}

func TestSaveAccount(t *testing.T) {
	err := refreshUserAndAccountTable()
	if err != nil {
		log.Fatalf("Error user and post refreshing table %v\n", err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}
	newAccount := models.Account{
		AccountNumber: "1111 2222 3333 4444",
		MoneyAmount:   74770,
		Comment:       "Test 0",
		UserID:        user.ID,
		Account:       user,
	}
	savedAccount, err := newAccount.SaveAccount(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the account: %v\n", err)
		return
	}
	assert.Equal(t, newAccount.ID, savedAccount.ID)
	assert.Equal(t, newAccount.AccountNumber, savedAccount.AccountNumber)
	assert.Equal(t, newAccount.MoneyAmount, savedAccount.MoneyAmount)
	assert.Equal(t, newAccount.UserID, savedAccount.UserID)
}

func TestDeleteAnAccount(t *testing.T) {

	err := refreshUserAndAccountTable()
	if err != nil {
		log.Fatalf("Error refreshing user and account table: %v\n", err)
	}
	account, err := seedOneUserAndOneAccount()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := accountInstance.DeleteAccount(server.DB, account.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
