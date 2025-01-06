package Day6

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Guard struct {
	x, y      int
	direction int
}

func loadGridFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

// Find the starting position and direction of the guard
func findStartingPoint(grid []string) Guard {
	var guard Guard
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				guard = Guard{x: i, y: j, direction: 0} // Start facing up
				break
			}
		}
	}
	return guard
}

func simulateMovement(grid []string, guard Guard) {

	steps := 0
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[string]bool) // Map to track unique visited positions

	// Simulate the guard's movement
	for {
		// Mark the current position as visited
		pos := fmt.Sprintf("%d,%d", guard.x, guard.y)
		if !visited[pos] {
			visited[pos] = true
			steps++
		}

		// Calculate the next position
		dx, dy := directions[guard.direction][0], directions[guard.direction][1]
		nx, ny := guard.x+dx, guard.y+dy

		// Check if the guard is stepping out of bounds
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
			break // Guard leaves the map
		}

		// Check if the next position is an obstacle
		if grid[nx][ny] == '#' {
			// Turn 90 degrees to the right
			guard.direction = (guard.direction + 1) % 4
		} else {
			// Move to the next position
			guard.x, guard.y = nx, ny
		}
	}
	fmt.Println("Steps taken:", steps)
}

func SixOne() { //Needs to start with Capital letter to be visible in another module
	grid, err := loadGridFromFile("Day6/input.txt")
	if err != nil {
		fmt.Println("Error loading map:", err)
		return
	}
	//fmt.Println(grid)

	guard := findStartingPoint(grid)
	fmt.Println(guard)
	simulateMovement(grid, guard)

}

func SixTwo() {
	grid, err := loadGridFromFile("Day6/input.txt")
	if err != nil {
		fmt.Println("Error loading map:", err)
		return
	}
	guard := findStartingPoint(grid)
	fmt.Println(guard)
	simulateMovement(grid, guard)

}
