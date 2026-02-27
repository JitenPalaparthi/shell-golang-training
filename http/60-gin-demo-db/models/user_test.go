package models_test

import (
	"http-demo/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserValidateSuccess(t *testing.T) {
	user := new(models.User)
	user.Id = 1
	user.Name = "Jiten"
	user.Email = "JitenP@Outlook.com"

	err := user.Validate()

	if err != nil {
		t.Fail() // text fail
		t.Log("Expected err is nil but actuall err is not nil")
	}
}

func TestUserValidateNameFailure(t *testing.T) {
	user := new(models.User)
	user.Id = 1
	user.Email = "JitenP@Outlook.com"

	err := user.Validate()
	if err == nil {
		t.Fail() // text fail
		t.Log("Expected err but returned as nil")
	}
}

func TestUserValidateEmailFailure(t *testing.T) {
	user := new(models.User)
	user.Name = "Jiten"
	user.Id = 1

	err := user.Validate()

	if err == nil {
		t.Fail() // text fail
		t.Log("Expected err but returned as nil")
	}
}

func TestUserValidateWrongEmailFailure(t *testing.T) {
	user := new(models.User)
	user.Id = 1
	user.Name = "Jiten"
	user.Email = "JitenPOutlook.com"
	err := user.Validate()

	if err == nil {
		t.Fail() // text fail
		t.Log("Expected err but returned as nil")
	}
}

func TestUserValidateSuccessR(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "jiten@gmail.com", false},
		{"missing @", "jitengmail.com", true},
		{"empty email", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &models.User{
				Id:    1,
				Name:  "Jiten",
				Email: tt.email,
			}

			err := u.Validate()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserValidateFailure(t *testing.T) {
	tests := []struct {
		name string
		user *models.User
		//wantErr bool
	}{
		{"invalid name", &models.User{Id: 101, Email: "Jitenp@outloo.com"}},
		{"invalid email", &models.User{Id: 101, Name: "Jiten"}},
		{"invalid email format", &models.User{Id: 101, Name: "Jiten", Email: "Jitenpoutloo.com"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			assert.Error(t, err)
		})
	}
}
