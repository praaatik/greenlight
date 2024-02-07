package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/praaatik/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("==> createMovieHandler")

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// fmt.Fprintf(w, "%+v\n", input)
	fmt.Println(input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("==> showMovieHandler")

	// moving all the logic for reading a parameter to a different helper function
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r, err)
		// http.NotFound(w, r)
		// return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Year:      2023,
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorReponse(w, r, err)
		// app.logger.Println(err)
		// http.Error(w, "The server encountered an error", http.StatusInternalServerError)
	}
}
