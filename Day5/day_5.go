package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read rules from file
	rules := make(map[int][]int)
	rulesFile, err := os.Open("Day5/rules.txt")
	if err != nil {
		fmt.Println("Error reading rules file:", err)
		return
	}
	defer rulesFile.Close()

	scanner := bufio.NewScanner(rulesFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		rules[from] = append(rules[from], to)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error parsing rules file:", err)
		return
	}

	// Read rows from file
	rowsFile, err := os.Open("Day5/rows.txt")
	if err != nil {
		fmt.Println("Error reading rows file:", err)
		return
	}
	defer rowsFile.Close()

	var rows [][]int
	scanner = bufio.NewScanner(rowsFile)
	for scanner.Scan() {
		line := scanner.Text()
		strNumbers := strings.Split(line, ",")
		var row []int
		for _, str := range strNumbers {
			num, _ := strconv.Atoi(str)
			row = append(row, num)
		}
		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error parsing rows file:", err)
		return
	}
	fmt.Println(rows)
	fmt.Println(rules)

	// Function to check if a row is valid
	isValidRow := func(row []int, rules map[int][]int) bool {
		for i := 0; i < len(row)-1; i++ {
			validFollowers, exists := rules[row[i]]
			if !exists {
				return false
			}
			valid := false
			for _, follower := range validFollowers {
				if row[i+1] == follower {
					valid = true
					break
				}
			}
			if !valid {
				return false
			}
		}
		return true
	}

	// Function to find the middle number of a row
	findMiddleNumber := func(row []int) int {
		return row[len(row)/2]
	}

	// Sum the middle numbers of valid rows
	validMiddleSum := 0
	numOfValidRows := 0
	for _, row := range rows {
		if isValidRow(row, rules) {
			validMiddleSum += findMiddleNumber(row)
			numOfValidRows++
		}
	}

	// Print the result
	fmt.Println("Sum of middle numbers of valid rows:", validMiddleSum, "Num of valid rows:", numOfValidRows)
}
