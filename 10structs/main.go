package main

import "fmt"

func main() {
	fmt.Println("Concept of structs")

	// no heritance in golang
	// no concept of super or parent

	divyam := User{"Divyam", "divyamraj1301@gmail.com", true, 22}

	// prints only values
	fmt.Println(divyam)

	// prints structures alongwith values
	fmt.Printf("Details of Divyam are : %+v\n", divyam)

	// accessing the fields using dot operator
	fmt.Printf("Name is : %+v\n", divyam.Name)
	fmt.Printf("Email is : %+v", divyam.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
