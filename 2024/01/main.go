package main

import (
	"bufio"
	"fmt"
	"os"
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
	right := make(map[int]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var l, r int

		_, err := fmt.Sscanf(line, "%d   %d", &l, &r)

		if err != nil {
			fmt.Printf("Coud not scan line: %s\n", line)
		}

		left = append(left, l)

		_, ok := right[r]
		if !ok {
			right[r] = 0
		}
		right[r]++
	}

	for _, v := range left {
		adder(v * right[v])
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
