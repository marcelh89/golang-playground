package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Page int `json:"page"`
	Data struct {
		Fruits []string `json:"fruits"`
	} `json:"data"`
}

func main() {

	jsonBody := `{"page": 1, "data" : { "fruits": ["apple", "peach"] }}`
	res := response{}
	json.Unmarshal([]byte(jsonBody), &res)
	fmt.Println(res)
}
