package controllers

import (
	"encoding/json"
	"goapi/src/authentication"
	"goapi/src/database"
	"goapi/src/models"
	"goapi/src/repositories"
	"goapi/src/responses"
	"goapi/src/security"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := func() ([]byte, error) {
		var r io.Reader = io.Reader(r.Body)
		b := make([]byte, 0, 512)
		for {
			n, err := r.Read(b[len(b):cap(b)])
			b = b[:len(b)+n]
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				return b, err
			}
			if len(b) == cap(b) {
				b = append(b, 0)[:len(b)]
			}
		}
	}()
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
	token, err := authentication.CreateToken(userStoredInDB.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}
