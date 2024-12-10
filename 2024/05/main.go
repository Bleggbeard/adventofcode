package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: something <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file '%s'\n", filename)
		os.Exit(2)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	adder := getAdder()
	ruleset := make(map[string][]string, 0)
	var comparator func(l, r string) int = nil
	for scanner.Scan() {
		line := scanner.Text()

		rule := strings.Split(line, "|")
		if len(rule) == 2 {
			processRule(rule, ruleset)
		} else {
			if comparator == nil {
				comparator = getComparator(ruleset)
			}
			manual := strings.Split(line, ",")
			if len(manual) > 1 {
				adder(processManual(manual, comparator))
			}
		}
	}

	fmt.Printf("Sum: %d\n", adder(0))
}

func processRule(rule []string, ruleset map[string][]string) {
	prev, after := rule[0], rule[1]
	_, ok := ruleset[prev]
	if !ok {
		ruleset[prev] = make([]string, 0)
	}
	ruleset[prev] = append(ruleset[prev], after)
}

func processManual(manual []string, comparator func(l, r string) int) int {
	orderedManual := slices.SortedFunc(slices.Values(manual), comparator)
	fmt.Printf("Manual:  %+v\n", manual)
	fmt.Printf("Ordered: %+v\n", orderedManual)
	if (isSameManual(manual, orderedManual)) {
		return 0
	}
	middle := orderedManual[(len(manual) - 1) / 2]
	ret, err := strconv.Atoi(middle)
	if err != nil {
		fmt.Printf("Could not parse %s\n", middle)
		return 0
	}
	return ret
}

func getComparator(ruleset map[string][]string) func(l, r string) int {
	return func(l, r string) int {
		if l == r {
			return 0
		}

		for _, b := range ruleset[l] {
			if r == b {
				return -1
			}
		}

		return 1
	}
}

func isSameManual(l, r []string) bool {
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}

	return true
}
