package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type MusicMax struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Name   string  `json:"name"`
	Artist *Author `json:"author"`
}

// slice is arrays слайс это вариабл а в эррейе нам придеться дать узындыгынаын значениясын
var musics []MusicMax

func getMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //сделали его стрингом типа соны кайтару керек деп
	json.NewEncoder(w).Encode(musics)
}

func getMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //сделали его стрингом типа соны кайтару керек деп
	params := mux.Vars(r)
	for _, item := range musics { // item = iterator
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&MusicMax{})
}

// to add new song to site
func createMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var music MusicMax
	_ = json.NewDecoder(r.Body).Decode(&music)
	music.ID = strconv.Itoa(rand.Intn(100)) //по сути нат зе бест чойс просто рандомно просто создает the id
	musics = append(musics, music)
	json.NewEncoder(w).Encode(music)
}

// we can обновить инфу
func updateMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range musics {
		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			var music MusicMax
			_ = json.NewDecoder(r.Body).Decode(&music)
			music.ID = params["id"]
			musics = append(musics, music)
			json.NewEncoder(w).Encode(music)
			return
		}
	}
}

func deleteMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range musics {
		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(musics)
}

func main() {

	router := mux.NewRouter()

	musics = append(musics, MusicMax{ID: "1", Isbn: "1234", Name: "Tadow", Artist: &Author{FirstName: "Masaego", LastName: "FKJ"}})
	musics = append(musics, MusicMax{ID: "2", Isbn: "5678", Name: "Далада", Artist: &Author{FirstName: "PRiNCE", LastName: "Папа"}})
	musics = append(musics, MusicMax{ID: "3", Isbn: "9012", Name: "35+34", Artist: &Author{FirstName: "Ariana", LastName: "Grande"}})

	router.HandleFunc("/api/musics", getMusics).Methods("GET") //это ссылканын сондары биз жасап жатырмыз
	router.HandleFunc("/api/musics/{id}", getMusic).Methods("GET")
	router.HandleFunc("/api/musics", createMusic).Methods("POST")
	router.HandleFunc("/api/musics/{id}", updateMusics).Methods("PUT")
	router.HandleFunc("/api/musics/{id}", deleteMusics).Methods("DELETE") // а тут можно айдига не только намберс но и стринги запихнуть можно
	//r.HandleFunc("/restaurants/{id:[0-9]+}", restaurant) прикол это типа онли фор диджитс

	// const PORT = ":8080" можно и так
	// log.Fatal(http.ListenAndServe(":8000", router)) BEFORE
	const port = ":8000"
	log.Fatal(http.ListenAndServe(port, router))
	log.Printf("starting server on %s \n", port)
	// 	log.Printf("Starting server on %s\n", PORT)

	errorr := http.ListenAndServe(port, router) //корейк не шыгады екеен, но пон что еррор хаха
	log.Fatal(errorr)                           // fatal это типа принтик, экуивелент

}
