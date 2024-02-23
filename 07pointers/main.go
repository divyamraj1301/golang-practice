package main

import "fmt"

func main() {
	fmt.Println("Concepts of Pointers")

	var ptr1 *int
	fmt.Println("Value of ptr1 is ", ptr1) // since no value has been assigned, value is nil

	myNumber := 23
	var ptr2 = &myNumber                        // assigning address of myNumber to ptr2
	fmt.Println("Value of myNumber is ", ptr2)  // actual memory address
	fmt.Println("Value of myNumber is ", *ptr2) // referenced value

	*ptr2 = *ptr2 + 2
	fmt.Println("New value is : ", myNumber)
}
