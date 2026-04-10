package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

	r := mux.NewRouter()

	//seeding
	courses = append(courses, []Course{
		{
			CourseId:   "CS101",
			CourseName: "Introduction to Programming",
			Price:      199,
			Author: &Author{
				FullName: "John Doe",
				Website:  "https://johndoe.dev",
			},
		},
		{
			CourseId:   "JS202",
			CourseName: "Advanced JavaScript",
			Price:      299,
			Author: &Author{
				FullName: "Jane Smith",
				Website:  "https://janesmith.io",
			},
		},
		{
			CourseId:   "GO303",
			CourseName: "Mastering Go",
			Price:      349,
			Author: &Author{
				FullName: "Abir Khan",
				Website:  "https://abirkhan.dev",
			},
		},
		{
			CourseId:   "DB404",
			CourseName: "Database Design",
			Price:      249,
			Author: &Author{
				FullName: "Michael Lee",
				Website:  "https://michaellee.dev",
			},
		},
	}...)


	
	//routing
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", craeteOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")






	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

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


func craeteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create a Course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Pls send some data")
	}

	//what about - {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if  course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func updateOneCourse (w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update a Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}


func deleteOneCourse (w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete a Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]... )
			break
		}
	}

	json.NewEncoder(w).Encode("Course has been deleted")

}