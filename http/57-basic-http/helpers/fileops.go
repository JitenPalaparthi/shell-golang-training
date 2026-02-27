package helpers

import (
	"encoding/json"
	"errors"
	"http-demo/models"
	"os"
)

func SaveUser(fn string, user *models.User) error {

	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	n, err := file.Write(bytes)
	if n == 0 {
		return errors.New("nothing to write")
	} else if err != nil {
		return err
	}

	return nil
}
