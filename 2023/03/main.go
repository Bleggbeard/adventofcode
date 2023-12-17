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
var symbolPositions [][]rune;

func parseLine(line string, row int) {
	num := ""
	startIndex := -1

	symbols := make([]rune, len(line))

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
				symbols[column] = char
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
	gearAdder := getAdder()
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		parseLine(line, row)

		row++
	}

	gearCandidates := make(map[string][]int)

	OuterLoop:
	for _, candidate := range partCandidates {
		minCol := int(math.Max(0, float64(candidate.topLeft.x)))
		minRow := int(math.Max(0, float64(candidate.topLeft.y)))
		maxCol := int(math.Min(float64(len(symbolPositions) - 1), float64(candidate.bottomRight.x)))
		maxRow := int(math.Min(float64(len(symbolPositions) - 1), float64(candidate.bottomRight.y)))

		for x := minCol; x <= maxCol; x++ {
			for y := minRow; y <= maxRow; y++ {
				if symbol := symbolPositions[y][x]; symbol != 0 {
					adder(candidate.value)
					if symbol == '*' {
						coords := fmt.Sprintf("%d.%d", y, x)
						gearCandidates[coords] = append(gearCandidates[coords], candidate.value)
					}
					continue OuterLoop
				}
			}
		}
	}

	for _, gears := range gearCandidates {
		if len(gears) == 2 {
			gearAdder(gears[0] * gears[1])
		}
	}

	fmt.Printf("Sum: %d\n", adder(0))
	fmt.Printf("Gear ratio sum: %d\n", gearAdder(0))
	// for row, symbolRow := range symbolPositions {
	// 	fmt.Printf("%d: ", row)
	// 	for _, symbol := range symbolRow {
	// 		if symbol == 0 {
	// 			fmt.Printf(" ")
	// 		} else {
	// 			fmt.Printf("%s", string(symbol))
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
}
