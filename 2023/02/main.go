package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

var gameRE = regexp.MustCompile("^Game (\\d*):")
var invalidRE = regexp.MustCompile("(?:1[3-9]|[2-9]\\d) red|(?:1[4-9]|[2-9]\\d) green|(?:1[5-9]|[2-9]\\d) blue")

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

		matches := gameRE.FindStringSubmatch(line)

		game := matches[1]
		gameNumber, err := strconv.Atoi(game)
		if err != nil {
			fmt.Printf("%s is invalid\n", line)
			continue
		}

		gameImpossible := invalidRE.MatchString(line)

		possibility := "possible"

		if gameImpossible {
			possibility = "not " + possibility
		}

		fmt.Printf("This is game #%d and it is %s\n", gameNumber, possibility)
		if !gameImpossible {
			adder(gameNumber)
		}
	}

	fmt.Printf("%d\n", adder(0))
}
