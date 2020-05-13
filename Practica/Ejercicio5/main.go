package main

import (
	"log"
	"net/http"

	"./peliculas"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", peliculas.ListaPeliculas).Methods("GET")
	router.HandleFunc("/pelicula/{id}", peliculas.GetPelicula).Methods("GET")
	router.HandleFunc("/pelicula/nueva", peliculas.NuevaPelicula).Methods("POST")
	router.HandleFunc("/pelicula/{id}", peliculas.ModificarPelicula).Methods("PUT")
	router.HandleFunc("/pelicula/{id}", peliculas.EliminarPelicula).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
