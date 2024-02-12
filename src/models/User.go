package models

import (
	"errors"
	"goapi/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// this represents a user of the social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Nameuser  string    `json:"nameuser,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pw        string    `json:"pw,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Prepare will call the methods validate and format to the added user
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Nameuser == "" {
		return errors.New("name is mandatory and cant be left blank")

	}
	if user.Nick == "" {
		return errors.New("nick is mandatory and cant be left blank")

	}
	if user.Email == "" {
		return errors.New("email is mandatory and cant be left blank")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("please insert a valid email")
	}

	if step == "signup" && user.Pw == "" {
		return errors.New("password is mandatory and cant be left blank")

	}
	return nil
}

func (user *User) format(step string) error {
	user.Nameuser = strings.TrimSpace(user.Nameuser)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "signup" {
		pwWithHash, err := security.Hash(user.Pw)
		if err != nil {

		}
		user.Pw = string(pwWithHash)
	}
	return nil
}
