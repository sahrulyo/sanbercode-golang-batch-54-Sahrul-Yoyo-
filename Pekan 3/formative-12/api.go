/*package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func Movies() []Movie {
	movs := []Movie{
		{1, "Spider-Man", 2002},
		{2, "John Wick", 2014},
		{3, "Avengers", 2018},
		{4, "Logan", 2017},
	}
	return movs
}

// GetMovies

func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/movies", getMovies)
	//post movie
	http.HandleFunc("/post_movie", PostMovie)
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// PostMovie
// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mov); err != nil {
				//new
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				log.Println("Error decoding JSON:", err)

			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			title := r.PostFormValue("title")
			getYear := r.PostFormValue("year")
			year, _ := strconv.Atoi(getYear)
			Mov = Movie{
				ID:    id,
				Title: title,
				Year:  year,
			}
		}

		dataMovie, _ := json.Marshal(Mov) // to byte
		w.Write(dataMovie)                // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

// pada function main
*/