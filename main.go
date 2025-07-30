package main

import (
	"fmt"
	"net/http"
	"time"

	application "github.com/VariableSan/go-http-crud/internal/app"
)

func main() {
	app, err := application.New()
	if err != nil {
		panic("")
	}

	app.Logger.Println("running app")

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
