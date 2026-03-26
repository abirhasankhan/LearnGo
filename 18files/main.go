package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Files in Go!")

	//write in file
	writeFile()

	readFile("./lcofile.text")

}

func writeFile() {

	content := "This is the content for file write in Go! :)"

	file, err := os.Create("lcofile.text")

	checkNilErr(err)

	len, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length of the file: ", len)

	defer file.Close()
}


func readFile(file string) {

	content, err := os.ReadFile(file)

	checkNilErr(err)

	fmt.Println("Content is: ", string(content))
}




func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
