package helpers

import (
	"encoding/json"
	"errors"
	"http-demo/models"
	"log/slog"
	"os"
)

var ChanUser chan *models.User

func init() { // no need to call
	ChanUser = make(chan *models.User, 100) // the buffe is suppose to unblock
	println("Some init")
	//go ProcessUsers()
}

func ProcessUsers(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		slog.Error(err.Error())
	}
	defer file.Close()

	for user := range ChanUser {

		bytes, err := json.Marshal(user)
		if err != nil {
			slog.Error(err.Error())
		}

		bytes = append(bytes, byte('\n'))

		n, err := file.Write(bytes)
		if n == 0 {
			slog.Error("nothing to write")
		} else if err != nil {
			slog.Error(err.Error())
		} else {
			slog.Info("user successfully stored")
		}
	}

}

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

	bytes = append(bytes, byte('\n'))

	n, err := file.Write(bytes)
	if n == 0 {
		return errors.New("nothing to write")
	} else if err != nil {
		return err
	}

	return nil
}
