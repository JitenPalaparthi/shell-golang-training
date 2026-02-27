package main

import (
	"flag"
	"http-demo/handlers"
	"http-demo/helpers"
	"log/slog"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func init() {
	PORT = os.Getenv("PORT")

}

func main() {

	//args := os.Args // --port=8089 -port=8089 --port 8089 -port 8089

	if PORT == "" {
		flag.StringVar(&PORT, "port", "8090", "--port=8089 -port=8089 --port 8089 -port 8089")
		flag.Parse()
	}

	router := gin.Default()

	commonHandler := handlers.NewCommonhandler()
	router.GET("/", commonHandler.Root)
	router.GET("/ping", commonHandler.Ping)
	router.GET("/health", commonHandler.Health("I am ok"))

	userHandler := handlers.NewUserHandler("users.db")

	go helpers.ProcessUsers(userHandler.FileName)

	user_router := router.Group("/v1/public")

	user_router.POST("/users", userHandler.Insert)

	if err := router.Run(":" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit() // before going to exit the main, it makes sure all other goroutines completed their execution
	} // all interfaces
}
