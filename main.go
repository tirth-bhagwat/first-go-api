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
	Director *Director `json:"author"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func main() {
	movies = append(movies,
		Movie{
			ID:    "123",
			Isbn:  "a123",
			Title: "mov-1",
			Director: &Director{
				FirstName: "Tirth",
				LastName:  "Bh",
			},
		},
	)

	r := mux.NewRouter()

	// r.HandleFunc("/api/movies", getMovies).Methods("GET")

	fmt.Println("Starting server at localhost:9090")

	r.HandleFunc("/getMovies", getMovies).Methods("GET")
	r.HandleFunc("/getMovie", getMovie).Methods("GET")
	r.HandleFunc("/", defaultRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", r))

}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is an API...</h1>")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, mov := range movies {
		if mov.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, mov := range movies {
		if mov.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[i])
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}

// func updateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var newMovie Movie
// 	json.NewDecoder(r.Body).Decode(&newMovie)

// }
