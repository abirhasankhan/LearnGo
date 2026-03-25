package main

import "fmt"

func main() {

	fmt.Println("Loop lesson")

	days := [] string{ "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	/* Method 1 */
	fmt.Println("--------------Method 1------------------")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	/* Method 2 */
	fmt.Println("--------------Method 2------------------")
	for i := range days {
		fmt.Println(days[i])
	}

	/* Method 3 */
	fmt.Println("--------------Method 3------------------")
	for index, day := range days {

		fmt.Printf("Index is %v and value is %v\n", index, day)
	}


	rougevalue := 1

	for rougevalue <= 10 {

		// if rougevalue == 5 {
		// 	break
		// }

		// if rougevalue == 4 {
		// 	rougevalue ++
		// 	continue
		// }

		if rougevalue == 7 {
			goto lco
		}

		fmt.Println("Value is: ", rougevalue)
		rougevalue ++


	}

	lco:
		fmt.Println("Jump from here !")



}
