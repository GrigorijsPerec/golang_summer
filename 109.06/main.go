package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const abc = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func main() {
	fmt.Printf("Эта программа подсчитывает, сколько ")
	fmt.Printf("различных маленьких русских букв в файле есть.\n\n")
	fmt.Printf("Выводит в файл количество букв в файле пример - а 156.\n")

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Чтение существующих значений из файла output.txt и загрузка в карту counter
	existingValues, err := readExistingValues(outputFile)
	if err != nil {
		fmt.Println("Ошибка при чтении существующих значений из файла:", err)
		return
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := make(map[rune]uint)
	n := 0

	// Загрузка существующих значений в карту counter
	for ch, count := range existingValues {
		counter[ch] = count
	}

	for scanner.Scan() {
		n++
		for _, ch := range scanner.Text() {
			ch = unicode.ToLower(ch)

			if ch == 'ё' {
				ch = 'е'
			} else if ch == 'ъ' {
				ch = 'ь'
			}

			if strings.ContainsRune(abc, ch) {
				counter[ch]++
			}
		}
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка при создании файла для вывода:", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	totalCount := float64(n)

	for ch, count := range counter {
		frequency := float64(count) / totalCount / 100
		line := fmt.Sprintf("%c %.2f\n", ch, frequency)
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Ошибка при сбросе буфера в файл:", err)
		return
	}

}

func readExistingValues(filename string) (map[rune]uint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	existingValues := make(map[rune]uint)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 2 {
			ch := rune(line[0])
			count := uint(0)
			existingValues[ch] = count
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return existingValues, nil
}
