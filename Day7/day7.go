package Day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// loadData reads the input file and returns the parsed target values and their associated numbers.
func loadData(filename string) ([]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targets := []int{}
	numberLists := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		// Parse the target value
		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, err
		}

		// Parse the numbers
		numStrs := strings.Fields(parts[1])
		nums := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			nums[i], err = strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}
		}

		targets = append(targets, target)
		numberLists = append(numberLists, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return targets, numberLists, nil
}

// evaluate checks if any combination of + and * operations between the numbers can result in the target value.
func evaluate(nums []int, target int) bool {
	n := len(nums)
	totalCombinations := 1 << (n - 1) // There are 2^(n-1) combinations of operators for n numbers

	for mask := 0; mask < totalCombinations; mask++ {
		result := nums[0]
		for i := 0; i < n-1; i++ {
			if (mask & (1 << i)) != 0 {
				// Apply + operator
				result += nums[i+1]
			} else {
				// Apply * operator
				result *= nums[i+1]
			}
		}
		if result == target {
			return true
		}
	}

	return false
}

// evaluateConcatenate checks if any combination of +, *, and || operations between the numbers can result in the target value.
func evaluateConcatenate(nums []int, target int) bool {
	n := len(nums)
	totalCombinations := 1
	for i := 0; i < n-1; i++ {
		totalCombinations *= 3 // There are 3^(n-1) combinations of operators for n numbers
	}

	for mask := 0; mask < totalCombinations; mask++ {
		result := nums[0]
		currentMask := mask
		for i := 0; i < n-1; i++ {
			operator := currentMask % 3 // Extract the operator for this pair
			currentMask /= 3

			switch operator {
			case 0:
				// Apply + operator
				result += nums[i+1]
			case 1:
				// Apply * operator
				result *= nums[i+1]
			case 2:
				// Apply || operator (concatenation)
				result = concatenate(result, nums[i+1])
			}
		}
		if result == target {
			return true
		}
	}

	return false
}

// concatenate combines two integers by treating them as strings and concatenating their digits.
func concatenate(a, b int) int {
	return toInt(fmt.Sprintf("%d%d", a, b))
}

// toInt converts a string to an integer.
func toInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func SevenOne() {
	targets, numberLists, err := loadData("Day7/input.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	totalSum := 0
	for i, target := range targets {
		if evaluate(numberLists[i], target) {
			totalSum += target
		}
	}

	fmt.Println("Total sum:", totalSum)
}

func SevenTwo() {
	targets, numberLists, err := loadData("Day7/input.txt")
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}

	totalSum := 0
	for i, target := range targets {
		if evaluateConcatenate(numberLists[i], target) {
			totalSum += target
		}
	}

	fmt.Println("Total sum:", totalSum)
}
