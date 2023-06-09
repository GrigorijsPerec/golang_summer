package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const down = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <input_file> <output_file> <shift>")
		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	shift := os.Args[3]

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("Error while opening the file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatal("Error while creating the output file:", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		encodedLine := encodeCaesar(line, shift)
		fmt.Fprint(writer, encodedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading the file:", err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal("Error while writing to the output file:", err)
	}

	fmt.Println("Encoding complete. The encoded output has been written to", outputFilePath)
}

func encodeCaesar(line string, shift string) string {
	shiftValue := parseShift(shift)
	encodedLine := ""

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			if strings.ContainsRune(down, ch) {
				encodedLine += string(encodeRune(ch, shiftValue, 'а', 'я'))
			} else if unicode.IsUpper(ch) {
				encodedLine += string(unicode.ToUpper(encodeRune(unicode.ToLower(ch), shiftValue, 'а', 'я')))
			}
		} else {
			encodedLine += string(ch)
		}
	}

	return encodedLine
}

func encodeRune(ch rune, shift int, start rune, end rune) rune {
	shiftedCh := (ch-start+rune(shift))%rune(end-start+1) + start
	return shiftedCh
}

func parseShift(shift string) int {
	shiftValue, err := strconv.Atoi(shift)
	if err != nil {
		log.Fatal("Invalid shift value:", err)
	}
	return shiftValue
}
