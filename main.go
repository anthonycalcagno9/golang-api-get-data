package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/anthonycalcagno9/golang-api-get-data/models"
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

	xmeals := []models.Meal{}

	json.Unmarshal(body, &xmeals)

	fmt.Println("after unmarshal = ", xmeals)

	for _, value := range xmeals {
		fmt.Println("Breakfast = ", value.Breakfast)
		fmt.Println("Lunch = ", value.Lunch)
		fmt.Println("Dinner = ", value.Dinner)
		fmt.Println("Rating = ", value.Rating)
		fmt.Println("DayOfTheWeek = ", value.DayOfTheWeek)
		fmt.Println("====================================")
	}

}
