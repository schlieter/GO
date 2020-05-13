package main

import "fmt"

type Usuario struct {
	edad             int
	nombre, apellido string
}

func main() {
	/*var user Usuario
	user.edad = 26
	user.nombre = "Cristian"
	user.apellido = "Schlieter"*/

	//user := Usuario{edad: 26, nombre: "Cristian", apellido: "Schlieter"}
	//fmt.Println(user.nombre)
	user := new(Usuario) //puntero a la estructura
	(*user).nombre = "Cristian"
	fmt.Println((*user)) // esto es medio ilegible por ende go te permite hacer esto fmt.Println(user)

}
