package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/src/authentication"
	"goapi/src/database"
	"goapi/src/models"
	"goapi/src/repositories"
	"goapi/src/responses"
	"goapi/src/security"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userStoredInDB, err := repository.SearchForEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPw(userStoredInDB.Pw, user.Pw); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	token, _ := authentication.CreateToken(userStoredInDB.ID)
	fmt.Println(token)
}
