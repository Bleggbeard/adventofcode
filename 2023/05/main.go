package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func getAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

const DST_INDEX = 0
const SRC_INDEX = 1
const RANGE_INDEX = 2

const MAP_SRC_INDEX = 1
const MAP_DST_INDEX = 2

var seedsRE = regexp.MustCompile("^seeds: (.*)$")
var mapRE = regexp.MustCompile("^([a-z]*)-to-([a-z]*) map:")

type mapping struct {
	source      string
	destination string
	mappings    conversionList
}

func (m *mapping) indexName() string {
	return fmt.Sprintf("%s", m.source)
}

type conversion struct {
	diff int
	start int
	end int
}

type conversionList struct {
	conversions []conversion
}

func (c conversionList) getOutput(i int) int {
	for _, conv := range c.conversions {
		if conv.start <= i && i <= conv.end {
			return i + conv.diff
		}
	}
	return i
}

func isMapping(line string) bool {
	return unicode.IsDigit([]rune(line)[0])
}

func getLocation(maps map[string]mapping, v int, m string) int {
	curMap, ok := maps[m]
	if !ok {
		return v
	}

	newValue := maps[m].mappings.getOutput(v)

	return getLocation(maps, newValue, curMap.destination)
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

	var seeds []string
	var currentMap *mapping
	almanacMappings := make(map[string]mapping)

	adder := getAdder()
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if isMapping(line) {
			if currentMap == nil {
				fmt.Printf("No map to assign mapping to?!\n")
				continue
			}

			mappingLine := strings.Split(line, " ")
			srcStart, srcErr := strconv.Atoi(mappingLine[SRC_INDEX])
			dstStart, dstErr := strconv.Atoi(mappingLine[DST_INDEX])
			mappingRange, rangeErr := strconv.Atoi(mappingLine[RANGE_INDEX])

			if srcErr != nil || dstErr != nil || rangeErr != nil {
				fmt.Printf("Could not parse mapping: %s\n", line)
				continue
			}

			conv := conversion{dstStart - srcStart, srcStart, srcStart + mappingRange}
			currentMap.mappings.conversions = append(currentMap.mappings.conversions, conv)

		} else if seedLine := seedsRE.FindStringSubmatch(line); seedLine != nil {
			seeds = strings.Split(seedLine[1], " ")
		} else if mapLine := mapRE.FindStringSubmatch(line); mapLine != nil {
			if currentMap != nil {
				almanacMappings[currentMap.indexName()] = *currentMap
			}

			src := mapLine[MAP_SRC_INDEX]
			dst := mapLine[MAP_DST_INDEX]

			currentMap = &mapping{
				src,
				dst,
				conversionList{make([]conversion, 0)},
			}
		}

		adder(len(line))
	}
	almanacMappings[currentMap.indexName()] = *currentMap

	for _, v := range almanacMappings {
		fmt.Printf("%s -> %s\n", v.source, v.destination)
	}

	minLocation := math.MaxInt
	for _, seed := range seeds {
		if s, err := strconv.Atoi(seed); err == nil {
			location := getLocation(almanacMappings, s, "seed")
			fmt.Printf("Seed #%d: %d\n", s, location)
			if location < minLocation {
				minLocation = location
			}
		}
	}

	fmt.Printf("Sum: %d\n", adder(0))
	fmt.Printf("Min location: %v\n", minLocation)
}
