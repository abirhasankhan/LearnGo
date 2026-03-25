package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to lesson on Slices!")

	
	// fruitList := []string{"Apple", "Tomato", "Mango"}

	// fmt.Printf("Type of slices: %T\n", fruitList)

	// fmt.Println(fruitList)

	// fruitList = append(fruitList, "Banana", "Orange")

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3])

	// fmt.Println(fruitList)


	// highScores := make([]int, 4)

	// highScores[0] = 10
	// highScores[1] = 54
	// highScores[2] = 36
	// highScores[3] = 88

	// fmt.Println(highScores)

	// // highScores[4] = 88 /*index out of range [4] with length 4*/

	// // fmt.Println(highScores)

	// highScores = append(highScores, 84, 65, 99)

	// fmt.Println(highScores)
	// fmt.Printf("Type of highScores: %T\n", highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	// sort.Ints(highScores)

	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))


	/* how to remove a value from slices based on index */

	var courses = []string{ "react", "js", "java", "swift", "python", "ruby"}

	fmt.Println("Courses: ", courses)

	var index int = 3

	courses = append(courses[:index], courses[index+1:]... )

	fmt.Println("Courses after deleting index nymber 3: ", courses)




}
