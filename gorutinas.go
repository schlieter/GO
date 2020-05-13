package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var stop string
	go mi_nombre_lentooo("Cristian")
	//Al agregar el "go" se transforma en una gorutina por ende corre detras del hilo principal(ASINCRONO)
	//codigo sincrono quiere decir que para que se
	//ejecute la siguiente linea tiene que terminar la que esta corriendo
	go func() {
		apellido_lento("schlieter")
	}()
	fmt.Println("que lentooo")

	fmt.Scanln(&stop)
}

func mi_nombre_lentooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
func apellido_lento(apellido string) {
	letras := strings.Split(apellido, "")
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
