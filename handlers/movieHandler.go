package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/movie-api/data"
	"github.com/movie-api/models"
	"github.com/movie-api/utils"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)


func GetMovies(w http.ResponseWriter, r *http.Request) {
	// set json header 
	// w.Header().Set("Content-Type", "application/json")
	// reuse it by putting it in response.go in utils folder
	utils.SetJSONHeader(w)
	json.NewEncoder(w).Encode(data.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	//set json header
	utils.SetJSONHeader(w)
	params := mux.Vars(r)

	for _, movie := range data.Movies{
		if movie.ID == params["id"]{
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// Header for not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No movie with this id")
}

func AddMovie(w http.ResponseWriter, r *http.Request){
	utils.SetJSONHeader(w)
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// movie.ID = strconv.Itoa(rand.Intn(1000000000)) 
	// replaced this by uuid to acoid potential collisions
	movie.ID = uuid.New().String()
	data.Movies = append(data.Movies, movie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Movie added successfully")
	json.NewEncoder(w).Encode(movie)
}

// updating with user provided data

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	utils.SetJSONHeader(w)
	params := mux.Vars(r)

	// deleting existing with the user provided id
	for index, movie := range data.Movies{
		if movie.ID == params["id"]{
			data.Movies = append(data.Movies[:index], data.Movies[index + 1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			data.Movies = append(data.Movies[:index], append([]models.Movie{movie}, data.Movies[index:]...)...)
			json.NewEncoder(w).Encode(movie)
            return
		}
	}
	// if not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("no movie for this id")
}

// Deleting movie by id provided by user

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	utils.SetJSONHeader(w)
	params := mux.Vars(r)

	for index,movie := range data.Movies{
		if movie.ID == params["id"]{
			data.Movies = append(data.Movies[:index], data.Movies[index + 1:]...)
			json.NewEncoder(w).Encode("Movie Deleted Successfully")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	//if not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("no movie for this id")
}