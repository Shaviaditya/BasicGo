package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	//Get Param
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// remove the old data
			movies = append(movies[:index], movies[index+1:]...)

			// Set a movie datatype
			var new_movie Movie

			//Use the pointer to store the respective data from the body
			_ = json.NewDecoder(r.Body).Decode(&new_movie)
			// Set its id
			new_movie.ID = params["id"]
			// append the new data
			movies = append(movies, new_movie)
			//Send the encoded JSON Data back
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "3000",
		Title: "Dune : Part 1",
		Director: &Director{
			Firstname: "Dennis",
			Lastname:  "Villenvue",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "3010",
		Title: "The Batman",
		Director: &Director{
			Firstname: "Matt",
			Lastname:  "Reeves",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		ISBN:  "3020",
		Title: "Dr Strange : Multiverse of Madness",
		Director: &Director{
			Firstname: "Sam",
			Lastname:  "Raimi",
		},
	})
	fileserve := http.FileServer(http.Dir("./static"))
	r.Handle("/",fileserve)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieId).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
