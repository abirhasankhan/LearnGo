package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in Golang")

	// var numOne int = 2
	// var numTwo float64 = 4

	// fmt.Println("Sum of the tow number is: ", numOne + int(numTwo))

	/*Random number*/
	// n := rand.Intn(5)
	// fmt.Println(n)

	/*Random from crypto*/

	num, _ := rand.Int(rand.Reader, big.NewInt(10))

	fmt.Println(num)


}