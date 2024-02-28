package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // aliases
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // slice of strings
}

func main() {
	fmt.Println("Create JSON data in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc123", []string{"webdev", "js"}},
		{"MERN Bootcamp", 99, "Youtube", "bcd123", []string{"mern", "full-stack"}},
		{"Angular Bootcamp", 199, "Youtube", "cde123", nil},
	}

	// package this data a JSON data
	finalJSON, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJson() {
	jsonFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "Youtube",
		"tags": ["webdev","js"]
	}
	`)
	var myCourses course

	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonFromWeb, &myCourses)
		fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("JSON is not valid")
	}

	// you want to add key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", key, val, val)
	}
}
