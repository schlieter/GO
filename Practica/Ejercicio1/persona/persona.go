package persona

import "fmt"

type Persona struct {
	Dni      int
	Nombre   string
	Apellido string
	Edad     int
}

type Habitantes []Persona

var persona *Persona
var habitantes Habitantes

func NuevaPersona() Persona {
	persona = new(Persona)
	fmt.Print("HOLA, ingrese sus datos.\nNombre: ")
	fmt.Scan(&persona.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&persona.Apellido)
	fmt.Print("Edad: ")
	fmt.Scan(&persona.Edad)
	fmt.Print("DNI: ")
	fmt.Scan(&persona.Dni)
	return *persona
}
func MostrarPersona(p Persona) {
	fmt.Println(p)
}

func AgregarHabitante(p Persona) Habitantes {
	habitantes = append(habitantes, p)
	return habitantes
}
func MostrarHabitantes(h Habitantes) {

	for _, habitantes := range h {
		fmt.Println(habitantes)
	}
}
