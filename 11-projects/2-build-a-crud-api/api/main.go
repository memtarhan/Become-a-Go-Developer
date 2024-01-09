package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	fmt.Println("Building an API for Movies")

	addMockupMovies()

	router := mux.NewRouter()
	setupRouters(router)

	fmt.Println("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func addMockupMovies() {
	movies = append(movies, Movie{
		ID:    1,
		ISBN:  "123",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "Director0 FirstName",
			LastName:  "Director0 LastName",
		},
	})

	movies = append(movies, Movie{
		ID:    2,
		ISBN:  "124",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Director1 FirstName",
			LastName:  "Director1 LastName",
		},
	})

	movies = append(movies, Movie{
		ID:    3,
		ISBN:  "125",
		Title: "Movie 3",
		Director: &Director{
			FirstName: "Director2 FirstName",
			LastName:  "Director2 LastName",
		},
	})
}

// - Setting up Routers
func setupRouters(router *mux.Router) {
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
}

// - Router Functions
func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	movieId, err := strconv.Atoi(parameters["id"])
	if err == nil {
		for _, movie := range movies {
			if movie.ID == movieId {
				json.NewEncoder(writer).Encode(movie)
				return
			}
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = rand.Intn(10000000)
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movies)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)

	parameters := mux.Vars(request)
	movieId, err := strconv.Atoi(parameters["id"])
	if err == nil {
		for index, item := range movies {
			if item.ID == movieId {
				movies = append(movies[:index], movies[index+1:]...)
				movie.ID = movieId
				movies = append(movies, movie)
				break
			}
		}
		json.NewEncoder(writer).Encode(movies)
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(request)
	movieId, err := strconv.Atoi(parameters["id"])
	if err == nil {
		for index, movie := range movies {
			if movie.ID == movieId {
				movies = append(movies[:index], movies[index+1:]...)
				break
			}
		}
		json.NewEncoder(writer).Encode(movies)
	}
}
