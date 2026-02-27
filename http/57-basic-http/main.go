package main

import (
	"flag"
	"http-demo/handlers"
	"log/slog"
	"net/http"
	"os"
	"runtime"
)

var (
	PORT string
)

func main() {

	//args := os.Args // --port=8089 -port=8089 --port 8089 -port 8089
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8090", "--port=8089 -port=8089 --port 8089 -port 8089")
		flag.Parse()
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	if r.Method != http.MethodGet {
	// 		w.WriteHeader(http.StatusNotImplemented)
	// 		return
	// 	}

	// 	w.Write([]byte("Hello Shell Folks"))
	// 	w.WriteHeader(http.StatusOK)
	// })

	commonHandler := handlers.NewCommonhandler()

	http.HandleFunc("/", commonHandler.Root)

	http.HandleFunc("/ping", commonHandler.Ping)
	http.HandleFunc("/health", commonHandler.Health("I am okay"))

	userHandler := handlers.NewUserHandler("users.db")

	http.HandleFunc("/users", userHandler.Insert)

	slog.Info("Server is running on->" + PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		slog.Error(err.Error())
		runtime.Goexit() // before going to exit the main, it makes sure all other goroutines completed their execution
	} // all interfaces
}
