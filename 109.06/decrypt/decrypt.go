package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadFreq(filename string) map[rune]int {
	frequency := make(map[rune]int)
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			letter := []rune(parts[0])[0]
			count := 0
			_, err := fmt.Sscanf(parts[1], "%d", &count)
			if err == nil {
				frequency[letter] = count
			}
		}
	}
	if scanner.Err() != nil {
		os.Exit(1)
	}
	return frequency
}

func findShift(ciphertext string, frequency map[rune]int) int {
	counts := make(map[int]int)
	for shift := 1; shift <= 32; shift++ {
		count := 0
		for _, char := range ciphertext {
			if char >= 'А' && char <= 'Я' {
				shifted := (int(char)-int('А')+shift+32)%32 + int('А')
				count += frequency[rune(shifted)]
			} else if char >= 'а' && char <= 'я' {
				shifted := (int(char)-int('а')+shift+32)%32 + int('а')
				count += frequency[rune(shifted)]
			}
		}
		counts[shift] = count
	}

	maxCount := 0
	bestShift := 0
	for shift, count := range counts {
		if count > maxCount {
			maxCount = count
			bestShift = shift
		}
	}

	return bestShift
}

func Decrypt(ciphertext string, shift int) string {
	plaintext := ""
	for _, char := range ciphertext {
		if char >= 'А' && char <= 'Я' {
			shifted := (int(char)-int('А')+shift+32)%32 + int('А')
			plaintext += string(rune(shifted))
		} else if char >= 'а' && char <= 'я' {
			shifted := (int(char)-int('а')+shift+32)%32 + int('а')
			plaintext += string(rune(shifted))
		} else {
			plaintext += string(char)
		}
	}
	return plaintext
}

func decryptFile(inputFile string, frequencyFile string, outputFile string) {
	frequency := loadFreq(frequencyFile)

	input, err := os.Open(inputFile)
	if err != nil {
		os.Exit(1)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		os.Exit(1)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ciphertext := scanner.Text()
		shift := findShift(ciphertext, frequency)
		plaintext := Decrypt(ciphertext, shift)
		fmt.Fprintln(output, plaintext)
	}
	if scanner.Err() != nil {
		os.Exit(1)
	}
}
func main() {
	if len(os.Args) != 4 {
		os.Exit(1)
	}
	inputFile := os.Args[1]
	freqFile := os.Args[2]
	outputFile := os.Args[3]

	decryptFile(inputFile, freqFile, outputFile)
}
