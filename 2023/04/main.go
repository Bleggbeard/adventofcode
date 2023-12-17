package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cardRE = regexp.MustCompile("(Card [ 0-9]*): ([ 0-9]*) \\| ([ 0-9]*)")

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

		matches := cardRE.FindStringSubmatch(line)
		winningNumbers := matches[2]
		chosenNumbers := matches[3]

		winning := strings.Split(winningNumbers, " ")
		chosen := strings.Split(chosenNumbers, " ")

		winningMap := make(map[int]bool)
		for _, winner := range winning {
			if winnerNumber, err := strconv.Atoi(winner); err == nil {
				winningMap[winnerNumber] = true
			}
		}

		points := 0
		for _, myNum := range chosen {
			if myNumber, err := strconv.Atoi(myNum); err == nil {
				if winningMap[myNumber] {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}

		adder(points)
	}

	fmt.Printf("Sum: %d\n", adder(0))
}
