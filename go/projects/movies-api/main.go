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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FullName string `json:"fullName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	populateMovies()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	fmt.Println("API is running at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func populateMovies() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "457685",
		Title: "The Founder",
		Director: &Director{
			FullName: "John Lee Hancock",
			LastName: "Lee Hancock",
		}})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "457684",
		Title: "The Wolf of Wall Street",
		Director: &Director{
			FullName: "Martin Scorsese",
			LastName: "Scorsese",
		}})
}

func getMovies(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	responseWriter.WriteHeader(http.StatusOK)

	json.NewEncoder(responseWriter).Encode(movies)
}

func deleteMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			responseWriter.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(responseWriter, "Movie not founded to delete", http.StatusNotFound)
}

func getMovieById(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range movies {
		if params["id"] == item.ID {
			json.NewEncoder(responseWriter).Encode(item)
			return
		}
	}

	http.Error(responseWriter, "Movie not found", http.StatusNotFound)

}

func createMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(100000)))

	movies = append(movies, movie)

	responseWriter.WriteHeader(http.StatusCreated)

	json.NewEncoder(responseWriter).Encode(movie)
}

func updateMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	var movie Movie

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(request.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)

			responseWriter.WriteHeader(http.StatusOK)

			json.NewEncoder(responseWriter).Encode(movie)

			return
		}
	}

	http.Error(responseWriter, "This movies does not exists", http.StatusNotFound)
}
