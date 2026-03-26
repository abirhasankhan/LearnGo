package main

import "fmt"

func main() {

	defer fmt.Println("Hello")

	fmt.Println("World")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("\nThree")

	demoDefer()



}


func demoDefer () {

	for i := 0; i < 5; i++ {

		defer fmt.Print(i)
	}
	
}
