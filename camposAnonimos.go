package main

import "fmt"

type Humano struct {
	nombre string
}

func (this Humano) hablar() string {
	return "Cristian"
}

type Animal struct {
	nombre string
}
type Tutor struct {
	Humano
	Animal
}

func (this Tutor) hablar() string { // de esta manera sobreescribo el metodo hablar de HUMANO
	return "Soy " + this.Humano.hablar()
}

/*
func (this Tutor) get_nombre() string {
	return this.nombre
}

func main() {
	persona := Tutor{Humano{nombre: "Cristian"}}
	fmt.Println(persona.get_nombre())
}
*/

func main() {
	persona := Tutor{Humano{nombre: "Cristian"}, Animal{nombre: "Luna"}}
	//Al tener dos campos anonimos que tienen un mismo atributo no sabe a cual se dirige
	//  por ende hay que indicar cual queremos mostrar
	fmt.Println(persona.Animal.nombre + " " + persona.Humano.hablar() + " " + persona.hablar())
}

//llamo al metodo de humano		//llamo al metodo de tutor
