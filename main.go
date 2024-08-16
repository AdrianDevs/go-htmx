package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"half.blue.gohtmx/handlers"
)

func main() {
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0) // Disable printing the time, source file, and line number.

	// names := []string{"Server", "Project", "Awesome"}
	// messages, err := greetings.Hellos(names)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(messages)

	// Logger
	// TODO

	// Service
	r := chi.NewRouter()

	//refresh := reload.New("views/")
	//r.Middlewares().Handler(refresh.Handle(r))

	// Ping
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", index)

	// user handlers
	userHandler := handlers.UserHandler{
		Service: "UserHandler service",
	}
	r.Get("/user", userHandler.Index)

	fmt.Printf("Starting server on port 8080\n")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("hiii"))
}
