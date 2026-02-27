package main

import (
	"flag"
	"http-demo/database"
	"http-demo/handlers"

	"log/slog"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	PORT, DB_URL string
)

func init() {
	PORT = os.Getenv("PORT")
	DB_URL = os.Getenv("DB_URL")
}

func main() {

	//args := os.Args // --port=8089 -port=8089 --port 8089 -port 8089

	if PORT == "" {
		flag.StringVar(&PORT, "port", "8090", "--port=8089 -port=8089 --port 8089 -port 8089")

	}
	if DB_URL == "" {
		flag.StringVar(&DB_URL, "db", `host=localhost user=postgres password=postgres dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "db=db string")
	}

	flag.Parse()
	router := gin.Default()

	commonHandler := handlers.NewCommonhandler()
	router.GET("/", commonHandler.Root)
	router.GET("/ping", commonHandler.Ping)
	router.GET("/health", commonHandler.Health("I am ok"))

	db, err := database.GetConnection(DB_URL)
	if err != nil {
		panic(err.Error())
	}

	IuserDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(IuserDB)

	GetUserRouterGroup(router, userHandler)

	if err := router.Run(":" + PORT); err != nil {
		slog.Error(err.Error())
		runtime.Goexit() // before going to exit the main, it makes sure all other goroutines completed their execution
	} // all interfaces
}

func GetUserRouterGroup(r *gin.Engine, userHandler *handlers.UserHandler) {
	user_router := r.Group("/v1/public")
	user_router.POST("/users", userHandler.Insert)
	user_router.GET("/users/:id", userHandler.GetUserBy)
	user_router.GET("/users/all", userHandler.GetAllUsers)
	user_router.DELETE("/users/:id", userHandler.GetDeleteBy)
}
