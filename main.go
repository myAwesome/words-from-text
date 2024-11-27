package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Використання: go run main.go input.txt output.txt")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Помилка при відкритті файлу %s: %v\n", inputFileName, err)
		return
	}
	defer file.Close()

	wordSet := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Зчитування слів
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		cleanedWord := cleanWord(word)
		if cleanedWord != "" {
			wordSet[cleanedWord] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Помилка при читанні файлу: %v\n", err)
		return
	}

	// Створення відсортованого списку слів без дублікатів
	uniqueWords := make([]string, 0, len(wordSet))
	for word := range wordSet {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Strings(uniqueWords)

	// Запис результату у вихідний файл
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Помилка при створенні файлу %s: %v\n", outputFileName, err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, word := range uniqueWords {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			fmt.Printf("Помилка при записі у файл: %v\n", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Помилка при закритті файлу: %v\n", err)
		return
	}

	fmt.Println("Файл успішно записаний:", outputFileName)
}

// Функція очищення слова від цифр і символів пунктуації
func cleanWord(word string) string {
	var cleaned strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(unicode.ToLower(r))
		}
	}
	return cleaned.String()
}
