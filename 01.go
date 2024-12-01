package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func advent01_1() {
	bytes, err := os.ReadFile("01.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")
	var a []int
	var b []int
	counter := 0

	for _, n := range data {
		line := strings.Split(n, "   ")

		first, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}
		a = append(a, first)
		second, err := strconv.Atoi(line[1])
		if err != nil {
			break
		}
		b = append(b, second)
	}

	sort.Ints(a)
	sort.Ints(b)

	// fmt.Println(a)
	// fmt.Println(b)

	for i := range a {
		if a[i] > b[i] {
			counter += a[i] - b[i]
		} else {
			counter += b[i] - a[i]
		}
	}

	fmt.Println(counter)
}

func advent01_2() {
	bytes, err := os.ReadFile("01.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")
	var a []int
	var b []int
	similarity := 0

	for _, n := range data {
		line := strings.Split(n, "   ")

		first, err := strconv.Atoi(line[0])
		if err != nil {
			break
		}
		a = append(a, first)
		second, err := strconv.Atoi(line[1])
		if err != nil {
			break
		}
		b = append(b, second)
	}

	for _, n := range a {
		counter := 0
		for _, m := range b {
			if n == m {
				counter += 1
			}
		}
		similarity += counter * n
	}

	fmt.Println(similarity)
}

func main() {
	advent01_1()
	advent01_2()
}
