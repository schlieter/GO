package main

import (
	"encoding/json"
	"net/http"

	"./persona"
)

func main() {
	p2 := persona.NuevaPersona()
	h := persona.AgregarHabitante(p2)

	p := persona.NuevaPersona()
	h = persona.AgregarHabitante(p)
	persona.MostrarHabitantes(h)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(h)
	})

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(p2)
	})
	http.ListenAndServe(":8080", nil)

}
