package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mikijunior/rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUserCreate(t *testing.T) {
	s := newServer(teststore.New())

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"email": "test@test.com",
				"password": "testPassword",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "empty email",
			payload: map[string]string{
				"email": "",
				"password": "testPassword",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "empty password",
			payload: map[string]string{
				"email": "email@email.com",
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
