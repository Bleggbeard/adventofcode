package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var mulRE = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

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
	for scanner.Scan() {
		line := scanner.Text()

		matches := mulRE.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])
			adder(l * r)
		}
	}

	fmt.Printf("Sum: %d\n", adder(0))
}
