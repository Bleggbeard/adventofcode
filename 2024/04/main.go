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
	left  Series
	up    Series
	right Series
}

type Series struct {
	running bool
	dir     rune
}

func (c *Character) toString() string {
	var sb strings.Builder
	sb.WriteRune(c.char)
	sb.WriteRune(c.stringifyBool(c.left.running, 'l'))
	sb.WriteRune(c.getDirOutput(c.left.dir))
	sb.WriteRune(c.stringifyBool(c.up.running, 'u'))
	sb.WriteRune(c.getDirOutput(c.up.dir))
	sb.WriteRune(c.stringifyBool(c.right.running, 'r'))
	sb.WriteRune(c.getDirOutput(c.right.dir))

	return sb.String()
}

func (c *Character) stringifyBool(v bool, r rune) rune {
	if v {
		return unicode.ToUpper(r)
	}
	return r
}

func (c *Character) getDirOutput(dir rune) rune {
	out := '.'

	switch dir {
	case 'f':
		out = 'v'
	case 'r':
		out = '^'
	}

	return out
}

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

var characterMatrix [][]Character = make([][]Character, 4)

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
		chars := make([]Character, len(line))

		hXmas := strings.Count(line, "XMAS")
		hSamx := strings.Count(line, "SAMX")
		adder(hXmas + hSamx)

		for col, char := range line {
			dir := '.'
			if char == 'X' {
				dir = 'f'
			} else if char == 'S' {
				dir = 'r'
			}

			left, up, right := Series{}, Series{}, Series{}
			if row > 0 {
				uc := characterMatrix[(row-1)%4][col]
				up.running = (uc.up.running || dir != '.') && isPreviousLetter(char, uc.char, uc.up.dir)

				if up.running && dir == '.' {
					up.dir = uc.up.dir
				}

				if row > 2 && isLastLetter(char, uc.up.dir) && up.running {
					adder(1)
					if uc.up.dir == 'f' {
						fmt.Printf("Found XMAS up at %d,%d\n", row, col)
					} else if uc.up.dir == 'r' {
						fmt.Printf("Found SAMX up at %d,%d\n", row, col)
					}
				}

				if col > 0 {
					lc := characterMatrix[(row-1)%4][col-1]
					left.running = (lc.left.running || dir != '.') && isPreviousLetter(char, lc.char, lc.left.dir)

					if left.running && dir == '.' {
						left.dir = lc.left.dir
					}

					if row > 2 && col > 2 && isLastLetter(char, lc.left.dir) && left.running {
						adder(1)
						if lc.left.dir == 'f' {
							fmt.Printf("Found XMAS left at %d,%d\n", row, col)
						} else if lc.left.dir == 'r' {
							fmt.Printf("Found SAMX left at %d,%d\n", row, col)
						}
					}
				}
				if col < len(line)-1 {
					rc := characterMatrix[(row-1)%4][col+1]
					right.running = (rc.right.running || dir != '.') && isPreviousLetter(char, rc.char, rc.right.dir)

					if right.running && dir == '.' {
						right.dir = rc.right.dir
					}


					if row > 2 && col < len(line)-3 && isLastLetter(char, rc.right.dir) && right.running {
						adder(1)
						if rc.right.dir == 'f' {
							fmt.Printf("Found XMAS right at %d,%d\n", row, col)
						} else if rc.right.dir == 'r' {
							fmt.Printf("Found SAMX right at %d,%d\n", row, col)
						}
					}
				}
			}

			if dir != '.' {
				left.dir = dir
				up.dir = dir
				right.dir = dir

				left.running = true
				up.running = true
				right.running = true
			}

			chars[col] = Character{char, left, up, right}
		}

		characterMatrix[row%4] = chars

		printMatrix()

		row++
	}

	fmt.Printf("Sum: %d\n", adder(0))
}

func isPreviousLetter(cur, prev, dir rune) bool {
	ret := false
	switch cur {
	case 'X':
		ret = prev == 'M' && dir == 'r'
	case 'M':
		ret = (prev == 'X' && dir == 'f') || (prev == 'A' && dir == 'r')
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
		return c == 'X'
	}
	return false
}

func printMatrix() {
	for i := range 4 {
		for _, v := range characterMatrix[i] {
			fmt.Printf("%s ", v.toString())
		}
		fmt.Println("")
	}

	fmt.Println("-----")
}
