package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(stones []int) []int {
	result := []int{}
	for _, stone := range stones {
		if stone == 0 {
			result = append(result, 1)
			continue
		}
		alphaStone := strconv.Itoa(stone)
		if len(alphaStone)%2 == 0 {
			alphaLeft := alphaStone[:len(alphaStone)/2]
			alphaRight := alphaStone[len(alphaStone)/2:]
			left, _ := strconv.Atoi(alphaLeft)
			right, _ := strconv.Atoi(alphaRight)
			result = append(result, left, right)
			continue
		}
		result = append(result, stone*2024)
	}
	return result
}

func advent11_1() {
	bytes, _ := os.ReadFile("11.txt")
	data := strings.Split(string(bytes[:len(bytes)-1]), " ")
	stones := []int{}
	for _, a := range data {
		i, _ := strconv.Atoi(a)
		stones = append(stones, i)
	}

	for range 25 {
		stones = blink(stones)
	}
	fmt.Println(len(stones))
}

func main() {
	advent11_1()
}
