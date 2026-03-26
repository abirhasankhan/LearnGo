package main

import "fmt"

func main() {

	fmt.Println("Lesson Structs")

	// no inheritance in golang; No super or parent or child

	Abir := User{"Abir Khan", "abir@gmail.com", true, 25}

	fmt.Println(Abir)

	fmt.Printf("User detail is %+v\n", Abir)

	fmt.Printf("User name is %v, and email is %v\n", Abir.Name, Abir.Email)

	Abir.GetStatus()
	Abir.NewMail()

	// fmt.Printf("User name is %v, and email is %v\n", Abir.Name, Abir.Email)


}


type User struct {
	Name string
	Email string
	Islogin bool
	Age int
}

func (u User) GetStatus() {

	fmt.Println("Is the user is login: ", u.Islogin)
}


func (u User) NewMail () {
	u.Email = "test@gmail.com"

	fmt.Println("The new mail is: ", u.Email)
}


