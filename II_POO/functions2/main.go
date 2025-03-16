package main

import "fmt"

// recibe un slice de valores enteros y devuelve un entero.
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

// recibe un slice de valores string y los imprime.
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Funcion que retorna multiples valores.
func getValues(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x // vieja forma.
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	// pasando datos en el argumento.
	fmt.Println(sum(3, 5, 7, 10, 25, 4))

	// Enviando datos en un arreglo.
	numbers := []int{4, 7, 3}
	fmt.Println(sum(numbers...))

	// Llamando a la funcion y pasando strings
	printNames("Carlos", "Jose", "Maria", "Laura")

	// Imprimiento multiples returns:
	fmt.Println(getValues(5))

	// capturando multiples returns:
	d, t, q := getValues(5)
	fmt.Println(d, t, q)
}
