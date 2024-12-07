package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

type Direction struct {
	dx int
	dy int
}

func advent06_1() {
	bytes, _ := os.ReadFile("06.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	height := len(lines)
	width := len(lines[0])

	// Make a grid, find the starting position
	pos := Position{0, 0}
	dir := Direction{0, 0}
	grid := make([][]rune, height)
	for i, line := range lines {
		grid[i] = make([]rune, width)
		for j, chr := range line {
			grid[i][j] = chr
			if chr == '^' {
				pos = Position{j, i}
				dir = Direction{0, -1}
			}
		}
	}

	fmt.Println("Starting position", pos)
	visited := make(map[Position]bool)
	visited[pos] = true
	for {
		// fmt.Println(pos, dir)
		if dir.dy == -1 { // Up
			if pos.y == 0 {
				break
			}
			if grid[pos.y-1][pos.x] == '#' {
				dir = Direction{dir.dx + 1, dir.dy + 1}
				continue
			}
			pos = Position{pos.x, pos.y - 1}
			visited[pos] = true
		} else if dir.dx == 1 { // Left
			if pos.x == width-1 {
				break
			}
			if grid[pos.y][pos.x+1] == '#' {
				dir = Direction{dir.dx - 1, dir.dy + 1}
				continue
			}
			pos = Position{pos.x + 1, pos.y}
			visited[pos] = true
		} else if dir.dy == 1 { // Down
			if pos.y == height-1 {
				break
			}
			if grid[pos.y+1][pos.x] == '#' {
				dir = Direction{dir.dx - 1, dir.dy - 1}
				continue
			}
			pos = Position{pos.x, pos.y + 1}
			visited[pos] = true
		} else if dir.dx == -1 { // Right
			if pos.x == 0 {
				break
			}
			if grid[pos.y][pos.x-1] == '#' {
				dir = Direction{dir.dx + 1, dir.dy - 1}
				continue
			}
			pos = Position{pos.x - 1, pos.y}
			visited[pos] = true
		}
	}

	fmt.Println(len(visited))
}

func main() {
	advent06_1()
}
