package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HeightMap struct {
	grid      [][]int
	trailends map[Point]bool
}

type Point struct {
	x int
	y int
}

func findNextSteps(grid [][]int, point Point) []Point {
	height := len(grid)
	width := len(grid[0])

	north := Point{point.x, point.y - 1}
	east := Point{point.x + 1, point.y}
	south := Point{point.x, point.y + 1}
	west := Point{point.x - 1, point.y}

	result := []Point{}
	if north.y >= 0 && grid[north.y][north.x]-grid[point.y][point.x] == 1 {
		result = append(result, north)
	}
	if east.x < width && grid[east.y][east.x]-grid[point.y][point.x] == 1 {
		result = append(result, east)
	}
	if south.y < height && grid[south.y][south.x]-grid[point.y][point.x] == 1 {
		result = append(result, south)
	}
	if west.x >= 0 && grid[west.y][west.x]-grid[point.y][point.x] == 1 {
		result = append(result, west)
	}
	return result
}

func markTrailends(heightMap HeightMap, point Point) {
	if heightMap.grid[point.y][point.x] == 9 {
		heightMap.trailends[point] = true
		return
	}

	nextSteps := findNextSteps(heightMap.grid, point)
	for _, step := range nextSteps {
		markTrailends(heightMap, step)
	}
}

func advent10_1() {
	bytes, _ := os.ReadFile("10.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	height := len(lines)
	width := len(lines[0])

	// Make a grid
	grid := make([][]int, height)
	for i, line := range lines {
		grid[i] = make([]int, width)
		for j, n := range line {
			num, _ := strconv.Atoi(string(n))
			grid[i][j] = num
		}
	}

	// For each 0 in the grid, count the trails and then sum unique ends with 9
	sum := 0
	for i, line := range grid {
		for j, cell := range line {
			if cell != 0 {
				continue
			}

			heightMap := HeightMap{grid, make(map[Point]bool)}
			markTrailends(heightMap, Point{j, i})
			sum += len(heightMap.trailends)

		}
	}
	fmt.Println(sum)
}

func main() {
	advent10_1()
}
