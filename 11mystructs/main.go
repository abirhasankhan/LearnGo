package main

import "fmt"

func main() {

	fmt.Println("Lesson Structs")

	// no inheritance in golang; No super or parent or child

	Abir := User{"Abir Khan", "abir@gmail.com", true, 25}

	fmt.Println(Abir)

	fmt.Printf("User detail is %+v\n", Abir)

	fmt.Printf("User name is %v, and email is %v", Abir.Name, Abir.Email)

}


type User struct {
	Name string
	Email string
	Islogin bool
	Age int
}


