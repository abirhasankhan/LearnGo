package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to user Pizza shop")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter yuor rating for the pizza: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Your rating is: ", input)
	fmt.Printf("Type of rating is: %T", input)

}
