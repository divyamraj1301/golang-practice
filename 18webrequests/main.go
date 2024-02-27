package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://divyamraj1301.netlify.app/"

func main() {
	fmt.Println("Web requests in golang")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is : %T\n", res)
	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
