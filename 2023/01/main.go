package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func findFirstDigit(line string, c chan string) {
	s := fwdReplacer.Replace(line)
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			c <- string(r[i])
			break
		}
	}

	c <- ""
}

func findLastDigit(line string, c chan string) {
	s := Reverse(line)
	revS := revReplacer.Replace(s)
	r := []rune(revS)
	for i := 0; i < len(r); i++ {
		if unicode.IsDigit(r[i]) {
			c <- string(r[i])
			break
		}
	}

	c <- ""
}

func Reverse(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

var fwdReplacer = strings.NewReplacer(
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9",
)

var revReplacer = strings.NewReplacer(
	Reverse("one"), "1",
	Reverse("two"), "2",
	Reverse("three"), "3",
	Reverse("four"), "4",
	Reverse("five"), "5",
	Reverse("six"), "6",
	Reverse("seven"), "7",
	Reverse("eight"), "8",
	Reverse("nine"), "9",
)

func convertNumbersToInts(line string) string {
	return fwdReplacer.Replace(line)
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
