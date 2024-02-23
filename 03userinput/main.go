package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// declaring buffered I/O
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating : ")

	// comma ok || error ok

	// taking input from user
	input, _ := reader.ReadString(('\n')) // no try cath over here, underscore used, or err can be used
	fmt.Println("Thanks for rating : ", input)
	fmt.Printf("Data type of rating : %T", input)
}
