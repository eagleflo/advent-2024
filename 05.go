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

func create_rules(rules_text []string) map[int]Rule {
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

func create_updates(updates_text []string) [][]int {
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

func update_in_order(rules map[int]Rule, update []int) bool {
	for i := range len(update) - 1 {
		before := update[:i]
		number := update[i]
		after := update[i+1:]
		order, _ := rules[number]
		for _, b := range before {
			if slices.Contains(order.after, b) {
				return false
			}
			b_order, _ := rules[b]
			if slices.Contains(b_order.before, number) {
				return false
			}
		}
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

func advent05_1() {
	bytes, _ := os.ReadFile("05.txt")
	data := strings.Split(string(bytes), "\n\n")
	rules_text := strings.Split(data[0], "\n")
	rules := create_rules(rules_text)
	updates_text := strings.Split(data[1], "\n")
	updates := create_updates(updates_text)

	sum := 0
	for _, update := range updates {
		if update_in_order(rules, update) {
			sum += update[len(update)/2]
		}
	}
	fmt.Println(sum)
}

func main() {
	advent05_1()
}
