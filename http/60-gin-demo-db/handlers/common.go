package handlers

import (
	"github.com/gin-gonic/gin"
)

func init() {
	println("some init in common handlers")
}

type CommonHandler struct {
}

func NewCommonhandler() *CommonHandler {
	return &CommonHandler{}
}

func (ch *CommonHandler) Root(ctx *gin.Context) {
	ctx.String(200, "Hello World!")
}

func (ch *CommonHandler) Ping(ctx *gin.Context) {

	//ctx.JSON(200, gin.H{"ping": "pong"}) // directly return json
	ctx.JSON(200, map[string]any{"ping": "pong"})

}

func (ch *CommonHandler) Health(msg string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{"health": msg})
	}
}
