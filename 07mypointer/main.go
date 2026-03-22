package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointer")

	// var ptr *int
	// fmt.Println("The value of the pointer: ", ptr)

	myNum := 23
	var ptr = &myNum

	fmt.Println("Memory location of the pointer: ", ptr)
	fmt.Println("The value of the pointer: ", *ptr)

	*ptr = *ptr - 3
	fmt.Println("New value: ", ptr)
	fmt.Println("New value: ", *ptr)



}