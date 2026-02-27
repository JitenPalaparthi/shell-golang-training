package handlers

import (
	"http-demo/helpers"
	"http-demo/models"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	FileName string
}

func NewUserHandler(fileName string) *UserHandler {
	return &UserHandler{fileName}
}

func (uh *UserHandler) Insert(ctx *gin.Context) {

	user := new(models.User)

	err := ctx.Bind(user)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	err = user.Validate()

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	user.Id = uint(rand.IntN(100000))
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	helpers.ChanUser <- user

	ctx.String(201, "user successfully stored in the file")

}
