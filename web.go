package main

import (
	"encoding/json"
	"net/http"
)

type Persona struct {
	Nombre   string //Al estar en minuscula la primer letra por default son atributos privados, lo mismo pasa con las func
	Apellido string
	edad     int
	dni      int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		persona := Persona{"Cristian", "Schlieter", 26, 37843888}
		json.NewEncoder(w).Encode(persona) //El primer parametro dice a donde vamos a enviar lo que pasemos a json
	})
	http.ListenAndServe(":8080", nil)
}
