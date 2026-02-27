package handlers

import (
	"encoding/json"
	"http-demo/helpers"
	"http-demo/models"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"
)

type UserHandler struct {
	FileName string
}

func NewUserHandler(fileName string) *UserHandler {
	return &UserHandler{fileName}
}

func (uh *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		slog.Error("unimplementedf header")
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte("invalid data"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := new(models.User)

	err = json.Unmarshal(data, user) // json deserialize

	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte("invalid data"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()

	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = uint(rand.IntN(100000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	err = helpers.SaveUser(uh.FileName, user)
	if err != nil {
		slog.Error(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("user successfully storefd"))
	w.WriteHeader(http.StatusCreated)

	// Store in a file

}
