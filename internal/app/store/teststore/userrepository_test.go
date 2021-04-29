package teststore_test

import (
	"testing"

	"github.com/mikijunior/rest-api/internal/app/model"
	"github.com/mikijunior/rest-api/internal/app/store"
	"github.com/mikijunior/rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(model.TestUser(t)))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	email := "test@email.com"

	s := teststore.New()
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
