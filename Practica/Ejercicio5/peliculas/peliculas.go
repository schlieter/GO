package peliculas

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Pelicula struct {
	Titulo   string `json:"titulo,omitempty"`
	Anio     int    `json:"anio,omitempty"`
	Idioma   string `json:"idioma,omitempty"`
	Sinopsis string `json:"sinopsis,omitempty"`
	ID       int    `json:"id,omitempty"`
}

var listaPeliculas []Pelicula

func NuevaPelicula(w http.ResponseWriter, r *http.Request) {
	var pelicula Pelicula
	//pelicula.ID = len(listaPeliculas) + 1
	_ = json.NewDecoder(r.Body).Decode(&pelicula)
	//listaPeliculas = append(listaPeliculas, pelicula)
	i, err := nuevaPeliculaDB(pelicula)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	//traer la pelicula nueva GetPeliculaDB(i)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetPeliculaDB(i))
}
func ListaPeliculas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ListaPeliculasDB())
}
func GetPelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Pelicula{})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetPeliculaDB(int64(id)))
}
func ModificarPelicula(w http.ResponseWriter, r *http.Request) {
	var pelicula Pelicula
	params := mux.Vars(r)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ListaPeliculasDB())
	}
	pelicula.ID = id
	json.NewDecoder(r.Body).Decode(&pelicula)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ModificarPeliculaDB(pelicula))

	/*
		for i, p := range listaPeliculas {
			if p.ID == id {
				json.NewDecoder(r.Body).Decode(&pelicula)
				pelicula.ID = p.ID
				buffer := append(listaPeliculas[:i], pelicula)
				listaPeliculas = append(buffer, listaPeliculas[i+1:]...)
				break
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listaPeliculas)*/

}
func EliminarPelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ListaPeliculasDB)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EliminarPeliculaDB(id))
}
