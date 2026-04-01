package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var baseurl string = "http://localhost:8000/"

func main() {

	fmt.Println("Web Request Server !")

	// PerformGetRequest(baseurl + "get")

	// PerformPostJSONRequest(baseurl + "post")

	PerformPostFormRequest(baseurl + "postform")
}

func PerformGetRequest(uri string) {

	res, err := http.Get(uri)

	if err != nil {

		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Lenght: ", res.ContentLength)

	content, _ := io.ReadAll(res.Body)

	// fmt.Println(string(content))

	var resString strings.Builder

	_, err = resString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println(resString.String())

}

func PerformPostJSONRequest(uri string) {

	// fake JSON payload

	reqestBody := strings.NewReader(`
		{
			"course": "Let's go with Golang",
			"price": 240,
			"platform":"n/a"

		}
	`)

	res, err := http.Post(uri, "application/json", reqestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}

func PerformPostFormRequest(uri string) {

	// form data
	data := url.Values{}

	data.Add("First Name", "Abir")
	data.Add("Last Name", "Khan")
	data.Add("Age", "26")

	res, err := http.PostForm(uri, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}
