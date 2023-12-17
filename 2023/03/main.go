package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

type position struct {
	x int
	y int
}

type schematicNumber struct {
	value       int
	topLeft     position
	bottomRight position
}

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func extractNumber(num string, topLeft position, bottomRight position) *schematicNumber {
	if num, err := strconv.Atoi(num); err == nil {
		return &schematicNumber{num, topLeft, bottomRight}
	}
	return nil
}

var partCandidates = make([]schematicNumber, 0)
var symbolPositions [][]bool;

func parseLine(line string, row int) {
	num := ""
	startIndex := -1

	symbols := make([]bool, len(line))

	for column, char := range []rune(line) {
		if unicode.IsDigit(char) {
			if num == "" {
				startIndex = column
			}
			num += string(char)
		} else {
			if num != "" {
				number := extractNumber(num, position{startIndex - 1, row - 1}, position{column, row + 1})
				if number != nil {
					partCandidates = append(partCandidates, *number)
				}
				num = ""
			}
			if char != '.' {
				symbols[column] = true
			}
		}
	}
	if num != "" {
		number := extractNumber(num, position{startIndex - 1, row - 1}, position{len(symbols) - 1, row + 1})
		if number != nil {
			partCandidates = append(partCandidates, *number)
		}
		num = ""
	}
	symbolPositions = append(symbolPositions, symbols)
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
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		parseLine(line, row)

		row++
	}

	OuterLoop:
	for _, candidate := range partCandidates {
		minCol := int(math.Max(0, float64(candidate.topLeft.x)))
		minRow := int(math.Max(0, float64(candidate.topLeft.y)))
		maxCol := int(math.Min(float64(len(symbolPositions) - 1), float64(candidate.bottomRight.x)))
		maxRow := int(math.Min(float64(len(symbolPositions) - 1), float64(candidate.bottomRight.y)))

		// fmt.Printf("Checking %d from {%d, %d} to {%d, %d}\n", candidate.value, minCol, minRow, maxCol, maxRow)
		for x := minCol; x <= maxCol; x++ {
			for y := minRow; y <= maxRow; y++ {
				// fmt.Printf("Checking [%d][%d]\n", y, x)
				if symbolPositions[y][x] {
					adder(candidate.value)
					fmt.Printf("%d (%v, %v) is a part number.\n", candidate.value, candidate.topLeft, candidate.bottomRight)
					continue OuterLoop
				}
			}
		}
		fmt.Printf("%d (%v, %v) is NOT a part number.\n", candidate.value, candidate.topLeft, candidate.bottomRight)
	}

	fmt.Printf("Sum: %d\n", adder(0))
	fmt.Printf("Rows: %d\n", row)
	// for y, symbolRow := range symbolPositions {
	// 	for x, symbol := range symbolRow {
	// 		if symbol {
	// 			fmt.Printf("[%d][%d]", x, y)
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
	// for row, symbolRow := range symbolPositions {
	// 	fmt.Printf("%d: ", row)
	// 	for _, symbol := range symbolRow {
	// 		if symbol {
	// 			fmt.Printf("x")
	// 		} else {
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
}
