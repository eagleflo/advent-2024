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

func advent08_1() {
	bytes, _ := os.ReadFile("08.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	grid := make([][]Location, len(lines))
	for i, line := range lines {
		grid[i] = make([]Location, len(line))
		for j, c := range line {
			grid[i][j] = Location{j, i, c, '.'}
		}
	}

	for i, line := range grid {
		for j, loc := range line {
			if loc.antenna != '.' {
				// Look ahead for all other antennae of this type
				// Calculate and mark their antinodes
				same_antennas := []Location{}
				// Scan line forward
				for _, pos := range line[j+1:] {
					if pos.antenna == loc.antenna {
						same_antennas = append(same_antennas, pos)
					}
				}
				// Scan rest of grid
				for _, line := range grid[i+1:] {
					for _, pos := range line {
						if pos.antenna == loc.antenna {
							same_antennas = append(same_antennas, pos)
						}
					}
				}
				for _, other := range same_antennas {
					dx := other.x - loc.x
					dy := other.y - loc.y

					nx := loc.x - dx
					ny := loc.y - dy
					if nx >= 0 && nx < len(grid[0]) &&
						ny >= 0 && ny < len(grid) {
						res := grid[ny][nx]
						grid[ny][nx] = Location{
							res.x,
							res.y,
							res.antenna,
							loc.antenna,
						}
					}

					mx := other.x + dx
					my := other.y + dy
					if mx >= 0 && mx < len(grid[0]) &&
						my >= 0 && my < len(grid) {
						res := grid[my][mx]
						grid[my][mx] = Location{
							res.x,
							res.y,
							res.antenna,
							loc.antenna,
						}
					}
				}
			}
		}
	}

	// Sum the number of antinodes on map
	sum := 0
	for _, line := range grid {
		for _, loc := range line {
			if loc.antinode != '.' {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}

func advent08_2() {

}

func main() {
	advent08_1()
}
