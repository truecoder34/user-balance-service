package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		errorMessage string
		name         string
		surname      string
		nickname     string
		email        string
		phone_number string
	}{
		{
			inputJSON:    `{"name": "Timothée1", "surname": "Chalame1", "nickname": "SweetBoy1", "email": "tc-sweet@gmail.com1", "phone_number": "+144712312341"}`,
			statusCode:   201,
			errorMessage: "",
			name:         "Timothée1",
			surname:      "Chalame1",
			nickname:     "SweetBoy1",
			email:        "tc-sweet@gmail.com1",
			phone_number: "+144712312341",
		},
		{
			inputJSON:    `{"name": "Timothée2", "surname": "Chalame2", "nickname": "SweetBoy2", "email": "tc-sweet@gmail.com1", "phone_number": "+144712312342"}`,
			statusCode:   500,
			errorMessage: "Email is already taken",
		},
		{
			inputJSON:    `{"name": "Timothée3", "surname": "Chalame3", "nickname": "SweetBoy1", "email": "tc-sweet@gmail.com2", "phone_number": "+144712312343"}`,
			statusCode:   500,
			errorMessage: "Nickname is already taken",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["nickname"], v.nickname)
			assert.Equal(t, responseMap["email"], v.email)
			assert.Equal(t, responseMap["name"], v.name)
			assert.Equal(t, responseMap["surname"], v.surname)
			assert.Equal(t, responseMap["phone_number"], v.phone_number)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
