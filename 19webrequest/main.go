package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://api.freeapi.app/api/v1/public/randomusers"



func main() {

	fmt.Println("Web Request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", res)

	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)

}
