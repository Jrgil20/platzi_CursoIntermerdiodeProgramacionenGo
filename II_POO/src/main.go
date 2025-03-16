package main

import "fmt"

type Employee struct {
	// se crea una estructura con dos campos privados
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	// se crea un metodo para modificar el campo privado id
	e.id = id
}

func (e *Employee) SetName(name string) {
	// se crea un metodo para modificar el campo privado name
	e.name = name
}

func (e *Employee) GetId() int {
	// se crea un metodo para obtener el campo privado id
	return e.id
}

func (e *Employee) GetName() string {
	// se crea un metodo para obtener el campo privado name
	return e.name
}

func main() {
	e := Employee{}
	//fmt.Printf("%v", e)
	e.id = 1
	e.name = "Name"
	//fmt.Printf("%v", e)
	e.SetId(5)
	e.SetName("Name 2")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
