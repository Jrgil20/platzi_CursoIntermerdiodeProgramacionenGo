package main

import "fmt"

func Generator(c chan<- int) {
	fmt.Println("in gen func-->")
	for i := 1; i <= 10; i++ {
		fmt.Println("for gen", i)
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	fmt.Println("in double func-->")
	for value := range in {
		fmt.Println("for double", value)
		out <- value * 2
	}
	close(out)
}

func Print(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
