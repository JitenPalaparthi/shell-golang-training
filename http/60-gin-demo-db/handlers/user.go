package handlers

import (
	"http-demo/database"
	"http-demo/models"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	database.IUserDB
}

func NewUserHandler(iuserdb database.IUserDB) *UserHandler {
	return &UserHandler{iuserdb}
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

	//user.Id = uint(rand.IntN(100000))
	// user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = uh.Create(user)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	ctx.JSON(201, user)

}

func (uh *UserHandler) GetUserBy(ctx *gin.Context) {

	_id, ok := ctx.Params.Get("id")

	if !ok {
		slog.Error("invalid id provided")
		ctx.String(http.StatusBadRequest, "invalid id provided")
		ctx.Abort()
	}

	id, err := strconv.Atoi(_id)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	user := new(models.User)

	user, err = uh.GetById(id)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	ctx.JSON(201, user)

}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := uh.GetAll()
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	ctx.JSON(201, users)
}

func (uh *UserHandler) GetDeleteBy(ctx *gin.Context) {

	_id, ok := ctx.Params.Get("id")

	if !ok {
		slog.Error("invalid id provided")
		ctx.String(http.StatusBadRequest, "invalid id provided")
		ctx.Abort()
	}

	id, err := strconv.Atoi(_id)

	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	r, err := uh.DeleteById(id)
	if err != nil {
		slog.Error(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}

	ctx.JSON(http.StatusNoContent, gin.H{"rows_affected": r})

}
