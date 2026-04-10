package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the course by Id")
	w.Header().Set("Content-Type", "application/json")

	parms := mux.Vars(r)

	for _, course := range courses {

		if course.CourseId == parms["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Now course found by that Id")
}