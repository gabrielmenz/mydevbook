package controllers

import (
	"encoding/json"
	"fmt"
	"goapi/src/database"
	"goapi/src/models"
	"goapi/src/repositories"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyrequest, err := func() ([]byte, error) {
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
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyrequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted: %d", userID)))

}
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
