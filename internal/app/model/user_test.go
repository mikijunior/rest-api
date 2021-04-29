package model_test

import (
	"testing"

	"github.com/mikijunior/rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.Validate())

	u.Email = "incorrect"

	assert.Error(t, u.Validate())
}
