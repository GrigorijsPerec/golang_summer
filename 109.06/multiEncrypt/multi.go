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
		fmt.Println("Usage: go run main.go <input_file> <output_file> <key_sequence>")
		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	keySequence := os.Args[3]

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
		encodedLine := encodeMultiCaesar(line, keySequence)
		fmt.Fprintln(writer, encodedLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading the file:", err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal("Error while writing to the output file:", err)
	}

	fmt.Println("Encoding complete. The encoded output has been written to", outputFilePath)
}

func encodeMultiCaesar(line string, keySequence string) string {
	encodedLine := ""
	keys := parseKeySequence(keySequence)
	keyIndex := 0

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			if strings.ContainsRune(down, ch) {
				encodedLine += string(encodeRune(ch, keys[keyIndex], 'а', 'я'))
				keyIndex = (keyIndex + 1) % len(keys)
			} else if unicode.IsUpper(ch) {
				encodedLine += string(unicode.ToUpper(encodeRune(unicode.ToLower(ch), keys[keyIndex], 'а', 'я')))
				keyIndex = (keyIndex + 1) % len(keys)
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

func parseKeySequence(keySequence string) []int {
	keys := strings.Split(keySequence, ",")
	keyValues := make([]int, len(keys))

	for i, key := range keys {
		keyValue, err := strconv.Atoi(strings.TrimSpace(key))
		if err != nil {
			log.Fatal("Invalid key sequence:", err)
		}
		keyValues[i] = keyValue
	}

	return keyValues
}
