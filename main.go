package main

// import all required modules
import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// define Movie struct
type Movie struct {
	ID string 				`json:"id"`
	Isbn string 			`json:"isbn"`
	Title string 			`json:"title"`
	Director *Director 		`json:"director"`
	// Pointer tells that Director is associated to Movie, it will be able to access fields that are present in Director struct
}

type Director struct{
	FirstName string		`json:"firstname"`
	LastName string 		`json:"lastname"`
}

// declaring a slice of type Movie struct to store all the movies data
var movies []Movie

// func to get all the movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	// setting content type to json so that struc data is converted to json or json to struct
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// id will be acceses through params
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			// deleting using append
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// decoding json data entered by user to type Movie struct in movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// creating random id for movie
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	// encoding movie to json to return to user to show that movie is created
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	// set json content type
	w.Header().Set("Content-Type","application/json")
	// getting acces to id using params
	params := mux.Vars(r)
	// looping over movies slice
	for index, item := range movies{
		if item.ID == params["id"]{
			// deleting the existing movie
			movies = append(movies[:index], movies[index + 1:]...)
			// adding the new movie with same id
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main(){
	// creating movie data
	movies = append(movies,Movie{
		ID:    "1",
		Isbn:  "1234567890",
		Title: "The Great Adventure",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	}, Movie{
		ID:    "2",
		Isbn:  "2345678901",
		Title: "Mystery at Dawn",
		Director: &Director{
			FirstName: "Jane",
			LastName:  "Smith",
		},
	}, Movie{
		ID:    "3",
		Isbn:  "3456789012",
		Title: "Sci-Fi Chronicles",
		Director: &Director{
			FirstName: "Albert",
			LastName:  "Einstein",
		},
	},Movie{
		ID:    "4",
		Isbn:  "4567890123",
		Title: "Romance in Paris",
		Director: &Director{
			FirstName: "Emily",
			LastName:  "Blunt",
		},
	},)

	// initialising a router using mux
	r := mux.NewRouter()
	// handling all routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}