package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/acalcagno9/golangapigetdata/models"
)

func main() {
	fmt.Println("Hello, World!")

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Printf("error with Get call = %v", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error Reading from response body: %v", err)
	}

	fmt.Println("body in bytes = ", body)

	fmt.Println("body as string = ", string(body))

	xmeals := models.Meal{}

}
