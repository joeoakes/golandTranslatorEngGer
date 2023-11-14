package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set your GPT-3 API key
	apiKey := ""

	// Create a request payload
	data := []byte(`{
        "prompt": "Translate the following English text to German: 'Hello, World!'",
        "max_tokens": 50,
        "temperature": 0.7
    }`)

	// Create an HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci/completions", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the API key in the request headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(responseBody))
}
