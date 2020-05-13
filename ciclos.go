package main

import "fmt"

/*func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hola mundo")
	}
}*/

/*func main() {
	i := 0
	for i < 10 {
		fmt.Println("hola mundo")
		i++
	}
}*/

func main() {
	i := 0
	for { //Ciclo infinito
		fmt.Println("hola mundo2")
		i++
		if i == 10 {
			break
		}
	}
}
