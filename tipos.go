package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 26
	edad_int, _ := strconv.Atoi("1") //El atoi devuelve dos valores , el convertido y otro de error con el _ se lo descarta
	fmt.Println("mi edad es ", strconv.Itoa(edad), " voy a cumplir ", edad+edad_int)
}
