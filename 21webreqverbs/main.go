package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Web Request Server !")

	PerformGetRequest("http://localhost:8000/get")
}

func PerformGetRequest(url string) {

	myurl := url

	res, err := http.Get(myurl)

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