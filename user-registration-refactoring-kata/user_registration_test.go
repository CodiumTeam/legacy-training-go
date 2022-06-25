package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_success_when_everything_is_valid(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPass_123123"
	req, _ := http.NewRequest(http.MethodPost, url, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestShould_returns_a_user_with_the_email_when_everything_is_valid(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPass_123123"
	req, _ := http.NewRequest(http.MethodPost, url, nil)

	router.ServeHTTP(w, req)

	var decodedResponse UserPostResponse
	json.Unmarshal(w.Body.Bytes(), &decodedResponse)
	assert.Equal(t, "my@Email.com", decodedResponse.Email)
}

func TestShould_returns_a_user_with_the_name_when_everything_is_valid(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPass_123123"
	req, _ := http.NewRequest(http.MethodPost, url, nil)

	router.ServeHTTP(w, req)

	var decodedResponse UserPostResponse
	json.Unmarshal(w.Body.Bytes(), &decodedResponse)
	assert.Equal(t, "Codium", decodedResponse.Name)
}

func TestShould_fail_when_password_is_short(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPas_1"
	req, _ := http.NewRequest(http.MethodPost, url, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Password is not valid", w.Body.String())
}

func TestShould_fail_when_password_does_not_contain_underscore(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPass123123"
	req, _ := http.NewRequest(http.MethodPost, url, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Password is not valid", w.Body.String())
}

func TestShould_fail_when_email_is_used(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	url := "/users?name=Codium&Email=my@Email.com&password=myPass_123123"
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	router.ServeHTTP(w, req)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "The email is already in use", w.Body.String())
}
