package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func make_grid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, v := range lines {
		numbers := strings.Split(v, " ")
		grid[i] = make([]int, len(numbers))
		for j, n := range numbers {
			number, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			grid[i][j] = number
		}
	}
	return grid
}

func safe(levels []int) bool {
	last_diff := 0
	for i, n := range levels[:len(levels)-1] {
		next := levels[i+1]
		diff := n - next
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if (last_diff < 0 && diff > 0) ||
			(last_diff > 0 && diff < 0) {
			return false
		}
		last_diff = diff
	}
	return true
}

func advent02_1() {
	bytes, err := os.ReadFile("02.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	grid := make_grid(lines)

	counter := 0
	for _, line := range grid {
		if safe(line) {
			counter += 1
		}
	}
	fmt.Println(counter)
}

func subsets(levels []int) [][]int {
	result := make([][]int, 0, len(levels))
	for i := range levels {
		before, after := levels[0:i], levels[i+1:]
		subset := make([]int, 0, len(levels)-1)
		subset = append(subset, before...)
		subset = append(subset, after...)
		result = append(result, subset)
	}
	return result
}

func advent02_2() {
	bytes, err := os.ReadFile("02.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	grid := make_grid(lines)

	counter := 0
OUTER:
	for _, line := range grid {
		// fmt.Println(line)
		if safe(line) {
			counter += 1
			continue
		}

		// Is there a safe subset?
		subsets := subsets(line)
		// fmt.Println(subsets)
		for _, s := range subsets {
			if safe(s) {
				counter += 1
				continue OUTER
			}
		}
	}
	fmt.Println(counter)
}

func main() {
	advent02_1()
	advent02_2()
}
