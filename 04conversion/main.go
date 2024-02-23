package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate us between 1 - 5")

	reader := bufio.NewReader(os.Stdin) // bufio  reader for reading from stdin
	input, _ := reader.ReadString('\n') // ReadString  reads until a newline is encountered. The returned data does not include the newline byte. 

	for {

	fmt.Println("Thanks for rating : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // strconv is used  to convert string to float64
	if err != nil {                                                    // strings is used  here because of the ReadString function
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating : ", numRating+1)
	}
}
