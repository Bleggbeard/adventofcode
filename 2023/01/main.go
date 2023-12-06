package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getAdder() func (int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func findFirstDigit(s string, c chan string) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			c <- string(r[i])
			break
		}
	}

	c <- ""
}

func findLastDigit(s string, c chan string) {
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsDigit(r[i]) {
			c <- string(r[i])
			break
		}
	}

	c <- ""
}

func main() {
	if (len(os.Args) != 2) {
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
		testString := scanner.Text()

		firstC := make(chan string)
		lastC := make(chan string)

		go findFirstDigit(testString, firstC)
		go findLastDigit(testString, lastC)

		first, last := <-firstC, <-lastC

		firstLast := fmt.Sprintf("%s%s", first, last)

		num, err := strconv.Atoi(firstLast)
		if err != nil {
			fmt.Printf("%s is invalid\n", testString)
			continue
		}

		adder(num)

	}
	fmt.Printf("%d\n", adder(0))
}
