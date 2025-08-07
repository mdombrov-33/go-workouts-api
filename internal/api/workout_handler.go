package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Workout ID: %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created a Workout\n")
}
