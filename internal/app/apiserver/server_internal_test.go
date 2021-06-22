package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/mikijunior/rest-api/internal/app/model"
	"github.com/mikijunior/rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"email":    "test@test.com",
				"password": "testPassword",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "empty email",
			payload: map[string]string{
				"email":    "",
				"password": "testPassword",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty password",
			payload: map[string]string{
				"email":    "email@email.com",
				"password": "",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}

			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/register", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleUserLogin(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email":    "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}

			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/login", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
