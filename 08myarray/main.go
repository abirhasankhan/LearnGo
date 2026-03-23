package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array Lesson")

	var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Mango"

	fmt.Println("The fruit list: ", fruitList)
	fmt.Println("The fruit list lenght", len(fruitList))

	var superHeros = [3] string{"Batman", "Ironman", "Superman"}

	fmt.Println("The fruit list: ", superHeros)
	fmt.Println("The fruit list lenght", len(superHeros))
}
