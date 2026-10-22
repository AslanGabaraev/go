package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go m.txt")
		os.Exit(1)
	}

	filePath := os.Args[1]

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}

	text := string(contentBytes)

	charCount := utf8.RuneCountInString(text)

	lineCount := strings.Count(text, "\n")
	if len(text) > 0 {
		lineCount++
	}

	words := strings.Fields(text)
	wordCount := len(words)

	fmt.Printf("\n--- Анализ файла: %s ---\n", filePath)
	fmt.Printf("Количество символов: %d\n", charCount)
	fmt.Printf("Количество слов: %d\n", wordCount)
	fmt.Printf("Количество строк: %d\n", lineCount)
}
