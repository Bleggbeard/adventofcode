package main

import (
	"bufio"
	"fmt"
	"os"
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
	for scanner.Scan() {
		line := scanner.Text()

		rule := strings.Split(line, "|")
		if len(rule) == 2 {
			processRule(rule, ruleset)
		} else {
			manual := strings.Split(line, ",")
			if len(manual) > 1 {
				adder(processManual(manual, ruleset))
			}
		}
	}
			fmt.Printf("Ruleset: %+v\n", ruleset)

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

func processManual(manual []string, ruleset map[string][]string) int {
	toCheck := make([]string, 0)
	fmt.Printf("Manual: %+v\n", manual)
	for _, entry := range manual {
		for _, c := range toCheck {
			for _, invalid := range ruleset[entry] {
				fmt.Printf("Entry: %s, C: %s, Inv: %s\n", entry, c, invalid)
				if c == invalid {
					return 0
				}
			}
		}
		toCheck = append(toCheck, entry)
		fmt.Printf("toCheck (%d): %+v\n", len(toCheck), toCheck)
	}
	middle := manual[(len(manual) - 1) / 2]
	ret, err := strconv.Atoi(middle)
	if err != nil {
		fmt.Printf("Could not parse %s\n", middle)
		return 0
	}
	return ret
}
