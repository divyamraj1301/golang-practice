package main

import "fmt"

func main() {
	fmt.Println("Concept of Maps")

	languages := make(map[string]string)

	// below is key || // below is related value
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS is abbreviated for: ", languages["JS"])

	//how to delete
	delete(languages, "RB")
	fmt.Println("After deleting Ruby from the map:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
