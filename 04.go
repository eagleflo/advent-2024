package main

import (
	"fmt"
	"os"
	"strings"
)

func advent04_1() {
	bytes, _ := os.ReadFile("04.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, c := range line {
			grid[i][j] = c
		}
	}

	sum := 0
	for i, line := range lines {
		for j, _ := range line {
			if grid[i][j] == 'X' {
				// Look in eight directions
				// EAST
				if j < len(line)-3 && grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S' {
					sum += 1
				}
				// SOUTHEAST
				if i < len(lines)-3 && j < len(line)-3 && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
					sum += 1
				}
				// SOUTH
				if i < len(lines)-3 && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
					sum += 1
				}
				// SOUTHWEST
				if i < len(lines)-3 && j >= 3 && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
					sum += 1
				}
				// WEST
				if j >= 3 && grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S' {
					sum += 1
				}
				// NORTHWEST
				if i >= 3 && j >= 3 && grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
					sum += 1
				}
				// NORTH
				if i >= 3 && grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
					sum += 1
				}
				// NORTHEAST
				if i >= 3 && j < len(line)-3 && grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
					sum += 1
				}
			}

		}
	}
	fmt.Println(sum)
}

func advent04_2() {
	bytes, _ := os.ReadFile("04.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, c := range line {
			grid[i][j] = c
		}
	}

	sum := 0
	for i, line := range lines {
		for j, _ := range line {
			if grid[i][j] == 'A' {
				if i < 1 || i > len(lines)-2 || j < 1 || j > len(line)-2 {
					continue
				}
				if ((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M') &&
					((grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') || grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M') {
					sum += 1
				}
			}

		}
	}
	fmt.Println(sum)
}

func main() {
	advent04_1()
	advent04_2()
}
