package main

import "fmt"

func main() {

	fmt.Println("Welcome to Map lesson")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["J"] = "Java"

	fmt.Println(languages)

	delete(languages, "PY")

	fmt.Println(languages)

	/*Loop*/

	for key, value := range languages {

		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
