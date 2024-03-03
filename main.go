package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually these are pointers , not values

func main() {
	// go greet("Hello")
	// greet("world")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.dev",
		"https://github.dev",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greet(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Status Code: %d, website: %s\n", res.StatusCode, endpoint)
	}

}
