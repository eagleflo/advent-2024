package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
	x        int
	y        int
	antenna  rune
	antinode rune
}

func prepareLocations() [][]*Location {
	bytes, _ := os.ReadFile("08.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	grid := make([][]*Location, len(lines))
	for i, line := range lines {
		grid[i] = make([]*Location, len(line))
		for j, c := range line {
			grid[i][j] = &Location{j, i, c, '.'}
		}
	}
	return grid
}

func findSameFreqAntennae(grid [][]*Location, i, j int, loc *Location) []*Location {
	result := []*Location{}
	// Scan line forward
	for _, pos := range grid[i][j+1:] {
		if pos.antenna == loc.antenna {
			result = append(result, pos)
		}
	}
	// Scan rest of grid
	for _, line := range grid[i+1:] {
		for _, pos := range line {
			if pos.antenna == loc.antenna {
				result = append(result, pos)
			}
		}
	}
	return result
}

func withinGrid(grid [][]*Location, x, y int) bool {
	return x >= 0 && x < len(grid[0]) &&
		y >= 0 && y < len(grid)
}

func sumAntinodes(grid [][]*Location) int {
	sum := 0
	for _, line := range grid {
		for _, loc := range line {
			if loc.antinode != '.' {
				sum += 1
			}
		}
	}
	return sum
}

func advent08_1() {
	grid := prepareLocations()
	for i, line := range grid {
		for j, loc := range line {
			if loc.antenna != '.' {
				// Calculate and mark their antinodes
				sameAntennae := findSameFreqAntennae(grid, i, j, loc)
				for _, other := range sameAntennae {
					dx, dy := other.x-loc.x, other.y-loc.y

					nx, ny := loc.x-dx, loc.y-dy
					if withinGrid(grid, nx, ny) {
						res := grid[ny][nx]
						res.antinode = loc.antenna
					}

					mx, my := other.x+dx, other.y+dy
					if withinGrid(grid, mx, my) {
						res := grid[my][mx]
						res.antinode = loc.antenna
					}
				}
			}
		}
	}
	fmt.Println(sumAntinodes(grid))
}

func advent08_2() {
	grid := prepareLocations()
	for i, line := range grid {
		for j, loc := range line {
			if loc.antenna != '.' {
				sameAntennae := findSameFreqAntennae(grid, i, j, loc)
				for _, other := range sameAntennae {
					dx, dy := other.x-loc.x, other.y-loc.y

					nx, ny := loc.x, loc.y
					for {
						if withinGrid(grid, nx, ny) {
							res := grid[ny][nx]
							res.antinode = loc.antenna
							nx -= dx
							ny -= dy
						} else {
							break
						}
					}

					mx, my := other.x, other.y
					for {
						if withinGrid(grid, mx, my) {
							res := grid[my][mx]
							res.antinode = loc.antenna
							mx += dx
							my += dy
						} else {
							break
						}
					}
				}
			}
		}
	}
	fmt.Println(sumAntinodes(grid))
}

func main() {
	advent08_1()
	advent08_2()
}
