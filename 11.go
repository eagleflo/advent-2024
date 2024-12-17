package main

import (
	"fmt"
	"math"
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
		// fmt.Println(stones)
	}
	fmt.Println(len(stones))
}

func numDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func splitInt(n int) (int, int) {
	divisor := int(math.Pow10(numDigits(n) / 2))
	left := n / divisor
	right := n % divisor
	return left, right
}

type Key struct {
	stone int
	count int
}

var cache = map[Key]int{}

func blinkOne(stone int, count int) int {
	if count == 0 {
		return 1
	}

	var result int
	key := Key{stone, count}
	cached, ok := cache[key]
	if ok {
		return cached
	}

	if stone == 0 {
		result = blinkOne(1, count-1)
	} else {
		stoneDigits := numDigits(stone)
		if stoneDigits%2 == 0 {
			left, right := splitInt(stone)
			result = blinkOne(left, count-1) + blinkOne(right, count-1)
		} else {
			result = blinkOne(stone*2024, count-1)
		}
	}
	cache[key] = result
	return result
}

func advent11_2() {
	bytes, _ := os.ReadFile("11.txt")
	data := strings.Split(string(bytes[:len(bytes)-1]), " ")
	stones := []int{}
	for _, a := range data {
		i, _ := strconv.Atoi(a)
		stones = append(stones, i)
	}

	sum := 0
	for _, stone := range stones {
		sum += blinkOne(stone, 75)
	}
	fmt.Println(sum)
}

func main() {
	advent11_1()
	advent11_2()
}
