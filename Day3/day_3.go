package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := "Day3/input_data.txt"

	//Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fileContent := string(data)
	//fmt.Println(fileContent)

	//Regex string
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(fileContent, -1)
	//fmt.Println(matches)

	totalSum := 0
	for _, match := range matches { // When ranging over slice for returns (index, copy of element)
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		totalSum += num1 * num2
	}
	fmt.Printf("Total Sum: %d\n\n", totalSum)

}
