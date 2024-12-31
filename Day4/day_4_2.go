package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define the file path
	filePath := "input_data.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read the grid from the file
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Dimensions of the grid
	rows := len(grid)
	cols := len(grid[0])

	// Function to check for MAS in X shape
	isXShape := func(x, y int) bool {
		// Check bounds for X shape
		if x-1 < 0 || x+1 >= rows || y-1 < 0 || y+1 >= cols {
			return false
		}

		if grid[x-1][y-1] == 'M' && grid[x][y] == 'A' && grid[x+1][y+1] == 'S' &&
			grid[x+1][y-1] == 'M' && grid[x-1][y+1] == 'S' {
			return true
		}

		if grid[x-1][y-1] == 'S' && grid[x][y] == 'A' && grid[x+1][y+1] == 'M' &&
			grid[x+1][y-1] == 'S' && grid[x-1][y+1] == 'M' {
			return true
		}

		if grid[x-1][y-1] == 'M' && grid[x][y] == 'A' && grid[x+1][y+1] == 'S' &&
			grid[x+1][y-1] == 'S' && grid[x-1][y+1] == 'M' {
			return true
		}

		if grid[x-1][y-1] == 'S' && grid[x][y] == 'A' && grid[x+1][y+1] == 'M' &&
			grid[x+1][y-1] == 'M' && grid[x-1][y+1] == 'S' {
			return true
		}

		return false
	}

	// Count occurrences
	count := 0

	// Iterate over each cell in the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if isXShape(x, y) {
				count++
			}
		}
	}

	// Print the total count
	fmt.Printf("Total occurrences of MAS in X shape: %d\n", count)
}
