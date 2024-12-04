package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	in, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Could not read input")
		os.Exit(3)
	}

	repl := strings.NewReplacer("do()", "\ndo()\n", "don't()", "\ndon't()\n")
	input := repl.Replace(string(in))

	inr := strings.NewReader(input)
	scanner := bufio.NewScanner(inr)

	adder := getAdder()
	mulEnabled := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "do()" {
			mulEnabled = true
		} else if line == "don't()" {
			mulEnabled = false
		}

		if mulEnabled {
			matches := mulRE.FindAllStringSubmatch(line, -1)

			for _, match := range matches {
				l, _ := strconv.Atoi(match[1])
				r, _ := strconv.Atoi(match[2])
				adder(l * r)
			}
		}
	}

	fmt.Printf("Sum: %d\n", adder(0))
}
