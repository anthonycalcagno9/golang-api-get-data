# Golang API Get Data

A simple Go application that demonstrates how to make HTTP GET requests to an API and parse JSON responses into Go structs.

## Overview

This application fetches meal data from a local API endpoint (`http://localhost:8080/`) and unmarshals the JSON response into a slice of `Meal` structs. It then displays the meal information in a formatted output.

## Features

- HTTP GET request to a local API
- JSON unmarshaling into Go structs
- Structured data display
- Error handling for network requests

## Project Structure

```
.
├── go.mod          # Go module definition
├── main.go         # Main application logic
├── models/
│   └── models.go   # Data models (Meal struct)
└── README.md       # This file
```

## Data Model

The application works with a `Meal` struct that contains:

- `Breakfast` (string) - Breakfast meal description
- `Lunch` (string) - Lunch meal description  
- `Dinner` (string) - Dinner meal description
- `Rating` (int) - Meal rating
- `DayOfTheWeek` (string) - Day of the week for the meal plan

## Prerequisites

- Go 1.23.3 or later
- A running API server on `localhost:8080` that returns meal data in JSON format

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/anthonycalcagno9/golang-api-get-data.git
   cd golang-api-get-data
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Ensure your API server is running on `localhost:8080` and returns meal data
2. Run the application:
   ```bash
   go run main.go
   ```

## Expected API Response Format

The API should return a JSON array of meal objects:

```json
[
  {
    "breakfast": "Scrambled eggs",
    "lunch": "Caesar salad",
    "dinner": "Grilled chicken",
    "rating": 8,
    "dayoftheweek": "Monday"
  }
]
```

## Output

The application will display:
- Raw response bytes
- Response as a string
- Parsed meal data with detailed breakdown for each day

## License

This project is open source and available under the [MIT License](LICENSE).