package main

import "fmt"

func main() {
	channel := make(chan string)

	go func(c chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			c <- name
		}
	}(channel)

	mensaje := <-channel
	fmt.Println(mensaje + " recibido")

	mensaje = <-channel
	fmt.Println(mensaje + " recibido2")
}
