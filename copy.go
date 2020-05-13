package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	//desti := make([]int, 3)
	desti := make([]int, len(slice), cap(slice)*2)
	copy(desti, slice) //si el destino es mas chico lo corta, para que esto no ocurra se recurre a len(variable origen) asi mantiene la longitud
	//otra practica comun es duplicar la capacidad
	fmt.Println(slice)
	fmt.Println(desti)

}
