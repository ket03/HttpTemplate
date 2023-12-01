package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	const URL = "https://reqres.in/" // editable
	const endPoint = "api/users?page=2" // editable
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL+endPoint, nil) // GET / POST
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	convertedBody := string(body)
	var data map[string]interface{}
	json.Unmarshal([]byte(convertedBody), &data)
	for key, value := range data {
		fmt.Println(key, value)
	}
}