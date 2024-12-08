package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent07_1() {
	bytes, _ := os.ReadFile("07.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]

	sum := 0
outer:
	for _, line := range lines {
		text := strings.Split(line, ": ")
		result, _ := strconv.Atoi(text[0])
		numbers := []int{}
		for _, s := range strings.Split(text[1], " ") {
			number, _ := strconv.Atoi(s)
			numbers = append(numbers, number)
		}

		// Interleave the operators between every pair
		calculations := []int{numbers[0]}
		for i := range len(numbers) - 1 {
			number := numbers[i+1]
			new_calculations := []int{}
			for _, c := range calculations {
				new_calculations = append(new_calculations, c+number, c*number)
			}
			calculations = new_calculations
		}

		for _, c := range calculations {
			if result == c {
				sum += result
				continue outer
			}
		}
	}

	fmt.Println(sum)
}

func advent07_2() {
	bytes, _ := os.ReadFile("07.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]

	sum := 0
outer:
	for _, line := range lines {
		text := strings.Split(line, ": ")
		result, _ := strconv.Atoi(text[0])
		numbers := []int{}
		for _, s := range strings.Split(text[1], " ") {
			number, _ := strconv.Atoi(s)
			numbers = append(numbers, number)
		}

		// Interleave the operators between every pair
		calculations := []int{numbers[0]}
		for i := range len(numbers) - 1 {
			number := numbers[i+1]
			new_calculations := []int{}
			for _, c := range calculations {
				c_str := strconv.Itoa(c)
				number_str := strconv.Itoa(number)
				concat, _ := strconv.Atoi(c_str + number_str)
				new_calculations = append(new_calculations, c+number, c*number, concat)
			}
			calculations = new_calculations
		}

		for _, c := range calculations {
			if result == c {
				sum += result
				continue outer
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	advent07_1()
	advent07_2()
}
