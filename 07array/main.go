package main

import "fmt"

func main() {
	fmt.Println("Concept of Arrays")

	var fruitArray [4]string // size of array has to be declared
	fruitArray[0] = "Apple"
	fruitArray[1] = "Banana"
	fruitArray[3] = "Orange"
	// fruitArray[2] is not initialized, it will give an extra space

	fmt.Println("Fruit Array: ", fruitArray)
	fmt.Println("Length: ", len(fruitArray)) // will return 4, couting the fruitArray[2] as blank space element

	var vegArray = [3]string{"potato", "mushroom", "tomato"}
	fmt.Println("Vegetable Array: ", vegArray)
	fmt.Println("Length: ", len(vegArray))
}
