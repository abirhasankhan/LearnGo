package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Plz rate our Pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating: ", input)
	fmt.Println("We are adding +1 to your rating: ")

	numRate, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("This is your raing after adding +1: ", numRate + 1)
	}



}