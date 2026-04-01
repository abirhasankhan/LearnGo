package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var baseurl string = "http://localhost:8000/"

func main() {

	fmt.Println("Web Request Server !")

	// PerformGetRequest(baseurl + "get")

	PerformPostJSONRequest(baseurl + "post")
}

func PerformGetRequest(url string) {

	res, err := http.Get(url)

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

func PerformPostJSONRequest(url string) {

	// fake JSON payload

	reqestBody := strings.NewReader(`
		{
			"course": "Let's go with Golang",
			"price": 240,
			"platform":"n/a"

		}
	`)

	res, err := http.Post(url, "application/json", reqestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}
