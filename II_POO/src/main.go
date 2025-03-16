package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

func main() {

	// Crear una instancia de Persona
	p := Persona{
		Nombre: "Juan",
		Edad:   30,
	}
	fmt.Printf("Nombre: %s, Edad: %d\n", p.Nombre, p.Edad)

	// Otra instancia de Persona
	p2 := Persona{
		Nombre: "Maria",
		Edad:   25,
	}
	fmt.Printf("Nombre: %s, Edad: %d\n", p2.Nombre, p2.Edad)
}
