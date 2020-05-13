package main

import "fmt"

type Usuario struct {
	edad             int
	nombre, apellido string
}

func (this Usuario) nombre_completo() string {
	return this.nombre + " " + this.apellido
}

func (this *Usuario) set_nombre(n string) { //PUNTERO A
	this.nombre = n
}

func main() {
	usuario := new(Usuario)
	usuario.nombre = "Cristian"
	usuario.apellido = "Schlieter"
	fmt.Println(usuario.nombre_completo())
	usuario.set_nombre("PEPE")
	fmt.Println(usuario.nombre_completo())
}
