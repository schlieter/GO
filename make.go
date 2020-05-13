package main

import "fmt"

func main() {
	slice := make([]int, 3, 3) //el make crea un slice pero inicializa en 0,"", o false dependiendo su tipo.El segundo elemento indica longitud. El tercer elemento indica la capacidad
	fmt.Println(cap(slice))
	slice = append(slice, 2) //el append agrega un elemento al final de la cadena si no tiene espacio lo redimenciona creando un nuevo slice
	fmt.Println(cap(slice))  // el cap muestra la capacidad de la cadena
}
