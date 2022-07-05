package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/truecoder34/user-balance-service/api/models"
	"github.com/truecoder34/user-balance-service/api/responses"
	"github.com/truecoder34/user-balance-service/api/utils/formaterror"
)

/*
	POST - create account
		[INPUT] - json body
*/
func (server *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account := models.Account{}
	err = json.Unmarshal(body, account)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account.Prepare()

	acCreated, err := account.SaveAccount(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, acCreated.ID))
	responses.JSON(w, http.StatusCreated, acCreated)
}

/*
	GET - get all accounts list handler
*/
func (server *Server) GetAccounts(w http.ResponseWriter, r *http.Request) {
	ac := models.Account{}
	acs, err := ac.FindAllAccounts(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, acs)
}

/*
	GET - Get account by its id
		[INPUT] - param ID
*/
func (server *Server) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := uuid.FromString(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	ac := models.Account{}
	acRecieved, err := ac.FindAccountByID(server.DB, aid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, acRecieved)

}

func (server *Server) UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

/*
	DELETE - delete account by id.
		[INPUT] - account id
		TODO: make it available only for admins
*/
func (server *Server) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid, err := uuid.FromString(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// check if account exists
	ac := models.Account{}
	err = server.DB.Debug().Model(models.Account{}).Where("id = ?", aid).Take(&ac).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	_, err = ac.DeleteAccount(server.DB, aid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%s", aid))
	responses.JSON(w, http.StatusNoContent, "")
}