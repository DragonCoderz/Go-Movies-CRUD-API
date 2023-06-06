package main

import (
	"fmt"
	"log"
	"encoding/json" //need this to encode data before sending it to postman
	"math/rand"
	"net/http" //how we create a server in go lang
	"strconv"
	"github.com/gorilla/mux"
)

 //Instead of using a fully fleshed database, for simplicity, we will be using structs and slices :)

 type Movie struct{
	ID string `json: "id"` //we do this json stuff so that we're able to encode and decode data when it comes through postman
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
 }

 type Director struct{
	Firstname string `json: ""`
	Lastname string `json: "lastname"`

 }

 var movies []Movie

 func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Essentially determines what type of stuff our w ResponseWriter is actually writing, which should be in json format
	json.NewEncoder(w).Encode(movies)
 }

 func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //This line uses the mux.Vars function from the Gorilla Mux package to retrieve the route parameters from the request. It extracts the parameters from the URL path, if any are specified. The parameters are stored in the params variable, which is a map of string keys to string values.
	for index,  item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
 }

 func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
 }

 func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
 }

 func updateMovie(w http.ResponseWriter, r* http.Request) {
	//Psuedo Code:
	//set json content type
	w.Header().Set("Content-Type", "applications/json")
	//get access to our params
	params := mux.Vars(r)
	//range over movies
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) //Returns movie to the user via postman
			return
		}
	}
	//delete the movie w/ the id that we've sent
	//then we'll just add a new movie that we sent in the body of postman



 }

 func main() {
	r := mux.NewRouter()

	//Populating our "database"
	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title: "Movie One", Director : &Director{Firstname: "John", Lastname:"Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn:"45455", Title: "Movie Two", Director : &Director{Firstname: "Steve", Lastname:"Smith"}})
	
	//The handle func assigns responsibilities for the inputted functions at their respective inputted directories  
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
 }