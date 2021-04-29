package store

import "github.com/mikijunior/rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(email string) (*model.User, error)
}
