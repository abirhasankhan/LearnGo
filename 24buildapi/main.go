package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"course"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}


var courses []Course 

func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Build API")
}


//controller 

func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to server home by Go!</h1>"))
}


func getAllCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All the course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}