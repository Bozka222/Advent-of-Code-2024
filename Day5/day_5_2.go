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
	//fmt.Println(rows)
	//fmt.Println(rules)

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

	// Function to reorder an invalid row based on the rules
	reorderRow := func(row []int, rules map[int][]int) []int {
		remaining := make([]int, len(row))
		copy(remaining, row) // copy of row with remaining values

		var reordered []int // new array for reordered row
		numCounter := 0

		fmt.Println("Reordered", reordered, "Remaining", remaining)
		// Continue until no numbers are left to place
		for len(remaining) > 0 {
			if len(reordered) == 0 {
				// Start with the first number
				reordered = append(reordered, remaining[numCounter])
				remaining = append(remaining[:numCounter], remaining[numCounter+1:]...)
			}

			fmt.Println("First number taken", "Reordered", reordered, "Remaining", remaining)
			current := reordered[len(reordered)-1]
			foundNext := false

			// Try to find the next valid number
			for i, num := range remaining {
				validFollowers, exists := rules[current]
				if exists {
					for _, validNext := range validFollowers {
						if num == validNext {
							reordered = append(reordered, num)
							remaining = append(remaining[:i], remaining[i+1:]...)
							foundNext = true
							break
						}
					}
				}
				if foundNext {
					break
				}
			}

			// If no valid continuation, try starting from another remaining number
			if !foundNext && len(remaining) > 0 {
				reordered = []int{}
				remaining = make([]int, len(row))
				copy(remaining, row) // copy of row with remaining values
				numCounter++

				fmt.Println("No valid continuation, Reordered", reordered, "Remaining", remaining)
			}
		}
		fmt.Println("Reordering done", "Reordered", reordered)
		return reordered
	}

	// Function to find the middle number of a row
	findMiddleNumber := func(row []int) int {
		return row[len(row)/2]
	}

	// Find invalid rows, reorder them, and calculate the middle sum
	invalidMiddleSum := 0
	numOfInvalidRows := 0
	for _, row := range rows {
		if !isValidRow(row, rules) {
			reordered := reorderRow(row, rules)
			invalidMiddleSum += findMiddleNumber(reordered)
			numOfInvalidRows++
		}
	}

	// Print the result
	fmt.Println("Sum of middle numbers of reordered invalid rows:", invalidMiddleSum, ", Num of invalid rows:", numOfInvalidRows)
}
