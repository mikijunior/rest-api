package model

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string `valid:"email,required"`
	Password          string `valid:"type(string),required,length(6|20)"`
	EncryptedPassword string
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)

	return err
}

func (u *User) BeforeCreate() error {
	if err := u.Validate(); err != nil {
		return err
	}

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
