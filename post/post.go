package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const URL = "https://reqres.in/" // editable
	const endPoint = "api/users" // editable
	client := &http.Client{}

	jsonFile, _ := os.Open("data.json")
	defer jsonFile.Close()
	body, _ := io.ReadAll(jsonFile)

	req, _ := http.NewRequest("POST", URL + endPoint, bytes.NewBuffer(body))
	res, _ := client.Do(req)
	text, _ := io.ReadAll(res.Body)
	convertedBody := string(text)
	var data map[string]interface{}
	json.Unmarshal([]byte(convertedBody), &data)
	for key, value := range data {
		fmt.Println(key, value)
	}
}