package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Concept of working with files in golang")
	content := "This will be input to a file"

	// to create a file
	file, err := os.Create("./myfile.txt")

	// if err != nil {
	// 	panic(err) // stop execution and will show the error
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) // write string into the file

	checkNilErr(err)
	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text in the file is : ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
