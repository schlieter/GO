package main

import "fmt"

func main() {
	/*slice := []int{12, 3, 55, 88} EL SLICE NO TIENE TAMAÑO PREDEFINIDO
	fmt.Println(slice)*/
	/*array := [3]int{1, 2, 3}
	slicing := array[0:2] //toma el array y lo redimenciona como indiquemos en los corchetes
	fmt.Println(slicing)*/
	/*
		array := make([]int, 3, 3)
		array = append(array, 1234, 2, 3, 4, 5)
		fmt.Println(array)
		for i := 0; i < len(array); i++ {
			buff := array
			if array[i] == 3 {
				//fmt.Println("hola", i, array[i])
				array = array[:i]
				buff = buff[i+1:]
				fmt.Println(buff)
			}
			array = append(array, buff[i])
		}
		fmt.Println(array)
	*/

	array := make([]int, 0, 0)
	buff := make([]int, 0, 0)
	array = append(array, 1234, 2, 3, 4, 5)
	for i := 0; i < len(array); i++ {

		if array[i] != 3 {
			buff = append(buff, array[i])

		}
	}
	fmt.Println(array)
	fmt.Println(buff)
}
