package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Reader(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error: Failed on Open Archive")
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		text := scanner.Text()

		number, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("linha Inválida", text)
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("erro ao ler arquivo:", err)
	}

	return numbers

}
