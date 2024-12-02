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
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, " ")
		safe := checkLevels(numbers)
		adder(safe)
	}

	fmt.Printf("Sum: %d\n", adder(0))
}

func checkLevels(numbers []string) int {
	prevNum := -1
	prevDir := -2

	for _, nums := range numbers {
		num, err := strconv.Atoi(nums)

		if err != nil {
			fmt.Printf("Could not parse number: %s", nums)
		}

		if prevNum != -1 {
			diff := num - prevNum
			absDiff := abs(diff)

			if absDiff < 1 || absDiff > 3 {
				return 0
			}

			dir := diff / absDiff
			if prevDir != -2 && dir != prevDir {
				return 0
			}

			prevDir = dir
		}

		prevNum = num
	}

	return 1
}

func abs(d int) int {
	if d < 0 {
		d = d * -1
	}

	return d
}
