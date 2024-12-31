package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
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
	//fmt.Println(grid)

	// Dimensions of the grid
	rows := len(grid)
	cols := len(grid[0])
	fmt.Printf("Rows: %d, Cols: %d\n", rows, cols)

	// Define the word to search
	word := "XMAS"
	wordLen := len(word)

	// Directions (dx, dy)
	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{0, -1},  // Left
		{-1, 0},  // Up
		{1, 1},   // Down-Right
		{1, -1},  // Down-Left
		{-1, 1},  // Up-Right
		{-1, -1}, // Up-Left
	}

	// Function to check if the word exists starting from a position
	isWordFound := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != word[i] {
				return false
			}
		}
		return true
	}

	// Iterate over each cell in the grid
	count := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if isWordFound(x, y, dx, dy) {
					count++
				}
			}
		}
	}

	// Print the results
	fmt.Printf("Total occurrences of '%s': %d\n", word, count)
}
