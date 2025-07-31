package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	application "github.com/VariableSan/go-http-crud/internal/app"
	"github.com/VariableSan/go-http-crud/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "server port")
	flag.Parse()

	app, err := application.New()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("server is running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
