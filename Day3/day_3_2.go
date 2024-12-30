package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := "input_data.txt"

	//Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fileContent := string(data)
	//fmt.Println(fileContent)

	//Regex string
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	instructions := regexp.MustCompile(`(?:(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)))`).FindAllString(fileContent, -1)

	totalSum := 0
	enabled := true
	for _, instr := range instructions { // When ranging over slice for returns (index, copy of element)
		switch {
		case doRegex.MatchString(instr):
			enabled = true
		case dontRegex.MatchString(instr):
			enabled = false
		case mulRegex.MatchString(instr) && enabled:
			match := mulRegex.FindStringSubmatch(instr)
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			totalSum += num1 * num2

		}
	}
	fmt.Printf("Total Sum: %d\n\n", totalSum)

}
