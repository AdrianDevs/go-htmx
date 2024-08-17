package main

import (
	"fmt"
	"half.blue.gohtmx/counter"
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

	countStore, err := count.NewStore("dummy db", "table name")
	if err != nil {
		log.Error("failed to create counter store", slog.Any("error", err))
	}
	countService := count.NewService(log, countStore)
	incHandler := count.NewHandler(log, countService)
	r.Get("/count", incHandler.Get)
	r.Post("/count", incHandler.Post)

	fmt.Printf("Starting server on port 8080\n")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
