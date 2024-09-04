package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var items []map[string]interface{}
	err = json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	chunkSize := 500
	totalItems := len(items)

	for i := 0; i < totalItems; i += chunkSize {
		end := i + chunkSize

		if end > totalItems {
			end = totalItems
		}

		chunk := items[i:end]
		fileName := fmt.Sprintf("./results/chunked-file-%d.json", i/chunkSize+1)
		chunkData, err := json.MarshalIndent(chunk, "", "  ")

		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		err = os.WriteFile(fileName, chunkData, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Printf("File %s written successfully\n", fileName)
	}
}
