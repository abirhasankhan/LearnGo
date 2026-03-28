package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://learn.go:5555/learn?coursename=golang&payment=sf654asf"

func main() {

	fmt.Println("Welcome to handling URLs in Go")

	result, err := url.Parse(myUrl)

	if err !=nil {
		panic(err)
	}

	fmt.Printf("Type of the url %T\n", result)


	// fmt.Println(result.Scheme)
	// fmt.Println(result.Path)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	query := result.Query()

	fmt.Println(query["coursename"])

	for i, val := range query {

		fmt.Printf("Key %v Value %v\n", i, val)
	}

	partOfUrl := &url.URL{
		Scheme: "https",
		Host: "learn.dev",
		Path: "/tutme",
		RawQuery: "user=Abir",
	}

	anotherUrl := partOfUrl.String()

	fmt.Println(anotherUrl)

}
