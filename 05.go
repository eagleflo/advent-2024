package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	before []int
	after  []int
}

func createRules(rules_text []string) map[int]Rule {
	// Make a rule mapping each number to 'before' and 'after' numbers
	rules := map[int]Rule{}
	for _, rule := range rules_text {
		split := strings.Split(rule, "|")
		before, _ := strconv.Atoi(split[0])
		after, _ := strconv.Atoi(split[1])
		before_rule, before_exists := rules[before]
		if !before_exists {
			rules[before] = Rule{[]int{}, []int{after}}
		} else {
			rules[before] = Rule{before_rule.before, append(before_rule.after, after)}
		}
		after_rule, after_exists := rules[after]
		if !after_exists {
			rules[after] = Rule{[]int{before}, []int{}}
		} else {
			rules[after] = Rule{append(after_rule.before, before), after_rule.after}
		}
	}
	return rules
}

func createUpdates(updates_text []string) [][]int {
	updates := [][]int{}
	for _, page := range updates_text[:len(updates_text)-1] {
		split := strings.Split(page, ",")
		numbers := []int{}
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			numbers = append(numbers, n)
		}
		updates = append(updates, numbers)
	}
	return updates
}

func updateIsInOrder(rules map[int]Rule, update []int) bool {
	for i := range len(update) - 1 {
		number := update[i]
		after := update[i+1:]
		order, _ := rules[number]
		for _, a := range after {
			if slices.Contains(order.before, a) {
				return false
			}
			a_order, _ := rules[a]
			if slices.Contains(a_order.after, number) {
				return false
			}
		}
	}
	return true
}

func fixUpdateOrder(rules map[int]Rule, update []int) []int {
	// fmt.Println("Fixing", update)
	for i := range len(update) - 1 {
		number := update[i]
		after := update[i+1:]
		order, _ := rules[number]
		for j, a := range after {
			if slices.Contains(order.before, a) {
				// A later number should be before this number, swap it here and continue
				new_update := slices.Clone(update)[:i]
				new_update = append(new_update, after[j])
				new_update = append(new_update, number)
				new_update = append(new_update, after[:j]...)
				new_update = append(new_update, after[j+1:]...)
				return fixUpdateOrder(rules, new_update)
			}
			a_order, _ := rules[a]
			if slices.Contains(a_order.after, number) {
				// A later number demands that this number is after them, swap and continue
				new_update := slices.Clone(update)[:i]
				new_update = append(new_update, after[:j+1]...)
				new_update = append(new_update, number)
				new_update = append(new_update, after[j+1:]...)
				return fixUpdateOrder(rules, new_update)
			}
		}
	}
	return update
}

func advent05_1() {
	bytes, _ := os.ReadFile("05.txt")
	data := strings.Split(string(bytes), "\n\n")
	rules_text := strings.Split(data[0], "\n")
	rules := createRules(rules_text)
	updates_text := strings.Split(data[1], "\n")
	updates := createUpdates(updates_text)

	sum := 0
	for _, update := range updates {
		if updateIsInOrder(rules, update) {
			sum += update[len(update)/2]
		}
	}
	fmt.Println(sum)
}

func advent05_2() {
	bytes, _ := os.ReadFile("05.txt")
	data := strings.Split(string(bytes), "\n\n")
	rules_text := strings.Split(data[0], "\n")
	rules := createRules(rules_text)
	updates_text := strings.Split(data[1], "\n")
	updates := createUpdates(updates_text)

	incorrectUpdates := [][]int{}
	for _, update := range updates {
		if !updateIsInOrder(rules, update) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	// Insight: there must be exactly one correct order for the result to add up
	sum := 0
	for _, update := range incorrectUpdates {
		correctUpdate := fixUpdateOrder(rules, update)
		sum += correctUpdate[len(correctUpdate)/2]
	}
	fmt.Println(sum)
}

func main() {
	advent05_1()
	advent05_2()
}
