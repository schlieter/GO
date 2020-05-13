package main

import "fmt"

type Profesionales interface {
	Datos() string
	//QueHago() string
}

/////////////////////////////////////////////////////////////////////////////////////////
ProductosInterface-datos,cantidadrestante,listado
Producto
Tipo-Producto

/////////////////
type Persona struct {
	nombre   string
	apellido string
	edad     int
	dni      int
}

/////////////////////////////////////////////////////////////////////////////////////////
type Programador struct {
	Persona
	profesion string
}

func (this Programador) Datos() string {
	return this.Persona.nombre + " " + this.Persona.apellido
}

/////////////////////////////////////////////////////////////////////////////////////////
type Contador struct {
	Persona
	profesion string
}

func (this Contador) Datos() string {
	return this.Persona.nombre + " " + this.Persona.apellido
}

/////////////////////////////////////////////////////////////////////////////////////////
func main() {
	prog1 := Programador{Persona{"Cristian", "Schlieter", 26, 37843888}, "programador"}
	profesionales := []Profesionales{prog1, Contador{Persona{"Yesica", "Chazarreta", 25, 38630980}, "contadora"}}

	for _, profesional := range profesionales {
		fmt.Println(profesional.Datos())
	}
}
