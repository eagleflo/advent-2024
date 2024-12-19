package main

import (
	"fmt"
	"os"
	_ "strconv"
	"strings"
)

func prepareTowelsAndDesigns() ([]string, []string) {
	bytes, _ := os.ReadFile("19.txt")
	parts := strings.Split(string(bytes), "\n\n")
	towels := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")
	designs = designs[:len(designs)-1]
	return towels, designs
}

func findFittingTowels(design string, towels []string) []string {
	var result []string
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			result = append(result, towel)
		}
	}
	return result
}

func findCompletion(design string, towels []string, candidates []string) int {
	// Figure out how far along the process we are already
	progress := 0
	for _, candidate := range candidates {
		progress += len(candidate)
	}

	// Check for base case, the completed design
	if progress == len(design) {
		return 1
	}

	// Next fitting towels
	fittingTowels := findFittingTowels(design[progress:], towels)
	for _, towel := range fittingTowels {
		result := findCompletion(design, towels, append(candidates, towel))
		if result == 1 {
			return 1
		}
	}

	return 0
}

func advent19_1() {
	towels, designs := prepareTowelsAndDesigns()

	sum := 0
	for _, design := range designs {
		sum += findCompletion(design, towels, []string{})
	}
	fmt.Println(sum)
}

func main() {
	advent19_1()
}
