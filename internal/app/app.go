package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mdombrov-33/go-api-fm/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New((os.Stdout), "", log.Ldate|log.Ltime)

	// our stores go here

	// our handlers go here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
