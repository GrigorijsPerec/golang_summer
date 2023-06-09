package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	const down = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

	inputF := os.Args[1]
	outputF := os.Args[2]

	file, err := os.Open(inputF)
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		return
	}
	defer file.Close()

	outputFile, err := os.Create(outputF)
	if err != nil {
		fmt.Println("Error while creating the output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)
	n := 0
	for scanner.Scan() {
		n++
		line := scanner.Text()
		convertedLine := ""
		for _, ch := range line {
			if strings.ContainsRune(down, ch) {
				convertedLine += string(ch)
			} else if unicode.IsUpper(ch) {
				lower := unicode.ToLower(ch)
				if strings.ContainsRune(down, lower) {
					convertedLine += string(lower)
				}
			}
		}
		fmt.Fprint(writer, convertedLine) // Изменено на fmt.Fprint
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error while writing to the output file:", err)
		return
	}

	fmt.Println("Conversion complete. The output has been written to output.txt.")
}
