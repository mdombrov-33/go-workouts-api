package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/mdombrov-33/go-api-fm/internal/app"
	"github.com/mdombrov-33/go-api-fm/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 4000, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Application started successfully on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatalf("Failed to start server: %v", err)
	}
}
