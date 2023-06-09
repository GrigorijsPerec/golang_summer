package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const down = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <input_file> <output_file> ")
		return
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

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
		encodedLine := encodeCaesar(line)
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

//func decrypt(){
//	make
//}

func bestSum