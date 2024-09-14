package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	count "half.blue.gohtmx/counter"
	"half.blue.gohtmx/fileserver"
	greet "half.blue.gohtmx/greetings"
	"half.blue.gohtmx/home"
	"half.blue.gohtmx/util"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	r := chi.NewRouter()

	// Ping
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(util.CheckIfHtmxMiddlerware)

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

	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Set FileServer parameters that will serve contents from the ./static/ folder.
	workDir, _ := os.Getwd()
	fileServer := fileserver.NewFileServer(log, http.Dir(filepath.Join(workDir, "static")))
	// FileServer routes files-server along /static
	fileServer.Handler("/static", r)

	fmt.Printf("Starting server on port 8080\n")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("Failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
