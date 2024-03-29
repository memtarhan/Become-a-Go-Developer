package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
	}

	app.logger.Println("id is: ", id)

	movie := models.Movie {
		ID: id,
		Title: "Some movie",
		Description: "Some description",
		Year: 2022,
		ReleaseDate: time.Date(2022, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime: 100,
		Rating: 5,
		MPAARating: "PG-13",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	
}