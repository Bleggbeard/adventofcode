package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Character struct {
	char  rune
	left  rune
	right rune
}

func (c *Character) toString() string {
	var sb strings.Builder
	sb.WriteRune(c.left)
	sb.WriteRune(c.char)
	sb.WriteRune(c.right)

	return sb.String()
}

func (c *Character) stringifyBool(v bool, r rune) rune {
	if v {
		return unicode.ToUpper(r)
	}
	return r
}

func getExpectedRune(i rune) rune {
	switch i {
	case 'M':
		return 'S'
	case 'S':
		return 'M'
	}

	return '.'
}

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

var characterMatrix [][]*Character = make([][]*Character, 3)

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
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		chars := make([]*Character, len(line))

		for col, char := range line {
			left, right, actualChar := '.', '.', char
			if row > 0 {
				if col > 0 && col < len(line)-1 && char == 'A' {
					pLeft := characterMatrix[(row-1)%3][col-1]
					pRight := characterMatrix[(row-1)%3][col+1]
					left = getExpectedRune(pRight.char)
					right = getExpectedRune(pLeft.char)

					if left == '.' || right == '.' {
						left = '.'
						right = '.'
					}
				}
				if col < len(line)-1 {
					pRight := characterMatrix[(row-1)%3][col+1]
					if pRight.left != char {
						pRight.right = '.'
					}
				}
				if col > 1 {
					pLeft := characterMatrix[(row-1)%3][col-1]
					if pLeft.right == char {
						adder(1)
						fmt.Printf("Found X-MAS at %d,%d\n", row-1, col-1)
					}
				}
			}

			chars[col] = &Character{actualChar, left, right}
		}

		characterMatrix[row%3] = chars

		printMatrix()

		row++
	}

	// 2035 is too high
	// 2023 is too high
	// 835 is too low
	fmt.Printf("Sum: %d\n", adder(0))
}

func isPreviousLetter(cur, prev, dir rune) bool {
	ret := false
	switch cur {
	case 'M':
		ret = (prev == 'A' && dir == 'r')
	case 'A':
		ret = (prev == 'M' && dir == 'f') || (prev == 'S' && dir == 'r')
	case 'S':
		ret = (prev == 'A' && dir == 'f')
	}

	return ret
}

func isLastLetter(c, dir rune) bool {
	switch dir {
	case 'f':
		return c == 'S'
	case 'r':
		return c == 'M'
	}
	return false
}

func printMatrix() {
	for i := range 3 {
		for _, v := range characterMatrix[i] {
			fmt.Printf("%s ", v.toString())
		}
		fmt.Println("")
	}

	fmt.Println("-----")
}
