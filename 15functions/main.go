package main

import "fmt"

func main() {

	fmt.Println("Welcome to FUNCTION lesson!")

	greeter()

	fmt.Println("Result: ", adder(5, 5))

	result, text := proAdder(5,4,5,54,554,9,4,6,4,1)

	fmt.Println("Pro sum: ", result)
	fmt.Println("Pro text: ", text)

}

func greeter () {
	fmt.Println("---Welcome here---")
}

func adder (valOne, valTwo int) int {

	return valOne + valTwo
}


func proAdder (values ... int) (int, string) {

	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi! I'm extra"
}