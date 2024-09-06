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

	const maxSize = 1 * 1024 * 1024 // 1.9 MB в байтах

	var currentChunk []map[string]interface{}
	currentSize := 0
	fileIndex := 1

	for _, item := range items {
		// Сериализуем отдельный элемент для оценки его размера
		itemData, err := json.Marshal(item)
		if err != nil {
			fmt.Println("Error marshaling item:", err)
			return
		}

		// Если добавление элемента приведет к превышению лимита, записываем текущий chunk
		if currentSize+len(itemData) > int(maxSize) {
			writeChunkToFile(currentChunk, fileIndex)
			fileIndex++
			currentChunk = nil
			currentSize = 0
		}

		// Добавляем элемент в текущий chunk и увеличиваем счетчик текущего размера
		currentChunk = append(currentChunk, item)
		currentSize += len(itemData)
	}

	// Записываем последний chunk, если остались не записанные элементы
	if len(currentChunk) > 0 {
		writeChunkToFile(currentChunk, fileIndex)
	}
}

// Функция для записи chunk в файл
func writeChunkToFile(chunk []map[string]interface{}, fileIndex int) {
	fileName := fmt.Sprintf("./results/chunked-file-%d.json", fileIndex)
	chunkData, err := json.MarshalIndent(chunk, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling chunk:", err)
		return
	}

	err = os.WriteFile(fileName, chunkData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("File %s written successfully\n", fileName)
}
