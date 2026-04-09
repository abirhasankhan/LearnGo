package main

import "fmt"


type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"course"`
	Price int `json:"price"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}


func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Build API")
}
