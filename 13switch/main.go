package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Switch Case")

	diceNum := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Value is 1 and you can open")
	case 2:
		fmt.Println("Value is 2 and you can move 2")
	case 3:
		fmt.Println("Value is 3 and you can move 3")
	case 4:
		fmt.Println("Value is 4 and you can move 4")
		fallthrough
	case 5:
		fmt.Println("Value is 5 and you can move 5")
	case 6:
		fmt.Println("Value is 6 and you can move 6 and roll the dice again!")
	default:
		fmt.Println("What was that!!")

	}
}
