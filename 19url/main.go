package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://divyamraj1301.netlify.app:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("Handling URLs in golang")
	fmt.Println(myurl)

	// parse URL
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)   // https
	fmt.Println(result.Host)     // divyamraj1301.netlify.app:3000
	fmt.Println(result.Path)     // learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs

	qparams := result.Query()
	fmt.Printf("Type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"]) // [reactjs]

	for _, val := range qparams {
		fmt.Println("Params is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "divyamraj1301.netlify.app",
		Path:    "/about-me",
		RawPath: "user=dev",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
