package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read json file
	filePath := "./input.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	escapedJSON, _ := json.Marshal(jsonData)
	result := strings.ReplaceAll(string(escapedJSON), "\"", "\\\"")

	fmt.Println(result)
}
