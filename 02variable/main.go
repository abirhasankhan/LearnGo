package main

import "fmt"


var name = "abir"

const LogInToken string = "asihdashdgasdhasd"

func main() {

	var usename string = "Abir"
	fmt.Println(usename)
	fmt.Printf("The type of this variable is %T \n", usename)

	var age int = 26
	fmt.Println(age)
	fmt.Printf("The type of this variable is %T \n", age)

	var isLogIn bool = true
	fmt.Println(isLogIn)
	fmt.Printf("The type of this variable is %T \n", isLogIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of this variable is %T \n", smallVal)

	var smallFloat float64 = 255.545411542524
	fmt.Println(smallFloat)
	fmt.Printf("The type of this variable is %T \n", smallFloat)


	var anotherVal float32
	fmt.Println(anotherVal)
	fmt.Printf("The type of this variable is %T \n", anotherVal)


	var numberOfUser = 3.0654
	fmt.Println(numberOfUser)
	fmt.Printf("The type of this variable is %T \n", numberOfUser)

	
	number := 3
	fmt.Println(number)
	fmt.Printf("The type of this variable is %T \n", number)

	fmt.Println(name)


	fmt.Println(LogInToken)
	fmt.Printf("The type of this variable is %T \n", LogInToken)

	


}
