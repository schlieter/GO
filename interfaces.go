package main

import "fmt"

type Usuario interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Nombre() string {
	return this.nombre
}

type Comun struct {
	nombre string
}

func (this Comun) Permisos() int {
	return 3
}
func (this Comun) Nombre() string {
	return this.nombre
}

func auth(usuario Usuario) string {
	permiso := usuario.Permisos()
	if permiso > 4 {
		return usuario.Nombre() + " tiene permiso de Administrador"
	} else if permiso <= 3 {
		return usuario.Nombre() + " tiene permiso Comun"
	}
	return ""
}

func main() {
	usuarios := []Usuario{Admin{nombre: "Cristian"}, Comun{nombre: "EPEPE"}} //Creo un array del tipo usuario(interface)
	for _, usuario := range usuarios {                                       //range devuelve un array con el indice y el "objeto"
		fmt.Println(auth(usuario)) // con el guion bajo descarto el indice y con for recorro el array
	}
}
