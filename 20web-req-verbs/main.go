package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web request verbs")
	// getRequest()
	// postJsonRequest()
	postFormRequest()
}

func getRequest() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code : ", res.StatusCode)
	fmt.Println("Content length : ", res.ContentLength)

	var resStr strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resStr.Write(content)
	fmt.Println("ByteCount : ", byteCount)
	fmt.Println(resStr.String())
}

func postJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	reqBody := strings.NewReader(`
			{
				"coursename": "Golang learning",
				"price": 250,
				"platform": "youtube"
			}
	`)

	res, err := http.Post(myUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func postFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("name", "divyam")
	data.Add("age", "22")
	data.Add("email", "divyam@go.dev")
	data.Add("country", "india")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
