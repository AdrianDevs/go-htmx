package main

import (
	"fmt"
	"half.blue.gohtmx/counter"
	greet "half.blue.gohtmx/greetings"
	"half.blue.gohtmx/home"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	r := chi.NewRouter()

	// Ping
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/", home.Index)

	greetStore, err := greet.NewStore("dummy greet db", "greet table name")
	if err != nil {
		log.Error("Failed to create greeting store", slog.Any("error", err))
	}
	greetService := greet.NewService(log, greetStore)
	greetHandler := greet.NewHandler(log, greetService)
	r.HandleFunc("/greet", greetHandler.ServeHTTP)

	countStore, err := count.NewStore("dummy count db", "count table name")
	if err != nil {
		log.Error("failed to create counter store", slog.Any("error", err))
	}
	countService := count.NewService(log, countStore)
	countHandler := count.NewHandler(log, countService)
	r.Get("/count", countHandler.Get)
	r.Post("/count", countHandler.Post)

	fmt.Printf("Starting server on port 8080\n")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
