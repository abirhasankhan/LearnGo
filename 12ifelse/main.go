package main

import "fmt"

func main() {

	fmt.Println("Lesson of If-else")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Oki NIGGA"
	} else {
		result = "error"
	}

	fmt.Println("Condition result: ", result)

	if 9%2 == 0 {
		fmt.Println("Result is even")
	} else {
		fmt.Println("Result is odd")
	}

	if num := 10; num < 10 {
		fmt.Println("Nymber is less then 10")
	} else {
		fmt.Println("Number is not less then 10")
	}


}
