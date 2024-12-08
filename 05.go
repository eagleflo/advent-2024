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

func correct_update_order(rules map[int]Rule, update []int) []int {
	// fmt.Println("Correcting", update)
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
				return correct_update_order(rules, new_update)
			}
			a_order, _ := rules[a]
			if slices.Contains(a_order.after, number) {
				// A later number demands that this number is after them, swap and continue
				new_update := slices.Clone(update)[:i]
				new_update = append(new_update, after[:j+1]...)
				new_update = append(new_update, number)
				new_update = append(new_update, after[j+1:]...)
				return correct_update_order(rules, new_update)
			}
		}
	}
	return update
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

func advent05_2() {
	bytes, _ := os.ReadFile("05.txt")
	data := strings.Split(string(bytes), "\n\n")
	rules_text := strings.Split(data[0], "\n")
	rules := create_rules(rules_text)
	updates_text := strings.Split(data[1], "\n")
	updates := create_updates(updates_text)

	incorrect_updates := [][]int{}
	for _, update := range updates {
		if !update_in_order(rules, update) {
			incorrect_updates = append(incorrect_updates, update)
		}
	}

	// Insight: there must be exactly one correct order for the result to add up
	sum := 0
	for _, update := range incorrect_updates {
		correct_update := correct_update_order(rules, update)
		sum += correct_update[len(correct_update)/2]
	}
	fmt.Println(sum)
}

func main() {
	advent05_1()
	advent05_2()
}
