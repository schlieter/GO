package main

import "fmt"

func main() {
	//var array [3]int
	array := [3]int{12, 56, 78}
	//fmt.Println(array)
	//fmt.Println(len(array)) //para saber el tamaño del array
	array[1] = 9
	fmt.Println(array)

	var matriz [2][3]int
	matriz[1][2] = 99
	fmt.Println(matriz)

}
