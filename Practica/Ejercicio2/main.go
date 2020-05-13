package main

import (
	"encoding/json"
	"net/http"

	"./lib"
)

//net/http

func main() {
	//	p := lib.NuevoProducto()
	//	p2 := lib.NuevoProducto()
	lib.AgregarProductoLista(lib.NuevoProducto())
	//	lib.AgregarProductoLista(p)
	//	lib.AgregarProductoLista(p2)
	lib.MostrarListaProducto()
	//lib.EliminarProducto()
	//lib.MostrarListaProducto()
	lib.ModificarProducto()
	lib.MostrarListaProducto()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(lib.MostrarJson())
	})
	http.ListenAndServe(":8080", nil)
}
