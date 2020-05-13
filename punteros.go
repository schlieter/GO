package main

import "fmt"

func main() {
	var puntero *int
	var punt *int
	h := 5
	puntero = &h //con esto retorno la direccion de la memoria
	h = 4
	punt = puntero //solo se puede asignar un puntero a otro puntero siendo del mismo tipo
	*punt = 10
	fmt.Println(*punt) //traigo el valor que esta en esa dir de mem
	fmt.Println(h)
}
