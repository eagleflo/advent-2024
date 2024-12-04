package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func sum_of_muls(batch string) int {
	mul_re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := mul_re.FindAllStringSubmatch(batch, -1)
	sum := 0
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		sum += first * second
	}
	return sum
}

func advent03_1() {
	bytes, _ := os.ReadFile("03.txt")
	data := string(bytes)
	fmt.Println(sum_of_muls(data))
}

func advent03_2() {
	bytes, _ := os.ReadFile("03.txt")
	data := string(bytes)
	dont_re := regexp.MustCompile(`don't\(\)`)
	do_re := regexp.MustCompile(`do\(\)`)
	batches := dont_re.Split(data, -1)
	sum := sum_of_muls(batches[0])
	for _, batch := range batches[1:] {
		split := do_re.Split(batch, -1)
		do_blocks := split[1:]
		for _, block := range do_blocks {
			sum += sum_of_muls(block)
		}
	}
	fmt.Println(sum)
}

func main() {
	advent03_1()
	advent03_2()
}
