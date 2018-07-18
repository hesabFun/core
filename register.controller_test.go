package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterNewUser(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var form struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	form.Mobile = "09111111114"
	form.Password = "12345678"
	form.Name = "test register new user"

	jsonValue, _ := json.Marshal(form)
	req, _ := http.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestRegisterNewUserFailed(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var form struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	form.Mobile = ""
	form.Password = ""
	form.Name = ""

	jsonValue, _ := json.Marshal(form)
	req, _ := http.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestVerifyUserBySms(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var form struct {
		SmsToken int    `json:"sms_token"`
		Mobile   string `json:"mobile"`
	}
	form.SmsToken = 1234
	form.Mobile = "09111111113"
	jsonValue, _ := json.Marshal(form)

	req, _ := http.NewRequest("POST", "/v1/auth/sms/verify", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestVerifyUserBySmsFailed(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var form struct {
		SmsToken int    `json:"sms_token"`
		Mobile   string `json:"mobile"`
	}
	form.SmsToken = 0
	form.Mobile = ""
	jsonValue, _ := json.Marshal(form)

	req, _ := http.NewRequest("POST", "/v1/auth/sms/verify", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestVerifyUserBySmsWrongCode(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	var form struct {
		SmsToken int    `json:"sms_token"`
		Mobile   string `json:"mobile"`
	}
	form.SmsToken = 4444
	form.Mobile = "09111111113"
	jsonValue, _ := json.Marshal(form)

	req, _ := http.NewRequest("POST", "/v1/auth/sms/verify", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
