package models

import (
	"errors"
	"regexp"
)

func init() { // no need to call

	println("Some init in user models")
}

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {

	if u.Name == "" {
		return errors.New("name is a must field")
	}

	if u.Email == "" {

		return errors.New("email is a must field")
	}

	if !validateEmail(u.Email) {
		return errors.New("invalid email")
	}
	// validation of email well

	return nil
}

func validateEmail(email string) bool {
	// RFC 5322–compliant simplified regex for most valid emails
	// This avoids overly complex patterns that can cause performance issues
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
