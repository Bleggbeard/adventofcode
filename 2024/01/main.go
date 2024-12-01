package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var l, r int

		_, err := fmt.Sscanf(line, "%d   %d", &l, &r)

		if err != nil {
			fmt.Printf("Coud not scan line: %s\n", line)
		}

		left = append(left, l)
		right = append(right, r)
		//TODO: Implement with binary tree or similar to increase performance
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		adder(abs(left[i], right[i]))
	}

	fmt.Printf("Sum: %d\n", adder(0))
}

func abs(l, r int) int {
	d := l - r

	if d < 0 {
		d = d * -1
	}

	return d
}
