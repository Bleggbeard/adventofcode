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

func getMultiplier() func(int) int {
	product := 1
	return func(i int) int {
		product *= i
		return product
	}
}

var whitespaceRE = regexp.MustCompile("\\s+")

func findFirstStrategy(t int, d int, c chan int) {
	for i := 1; i < t; i++ {
		if i*(t-i) > d {
			c <- i
			return
		}
	}
	c <- -1
}

func findLastStrategy(t int, d int, c chan int) {
	for i := t - 1; i > 0; i-- {
		if i*(t-i) > d {
			c <- i
			return
		}
	}
	c <- -1
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

	var times []int
	var distances []int

	adder := getAdder()
	multiplier := getMultiplier()
	for scanner.Scan() {
		line := scanner.Text()

		lineData := whitespaceRE.Split(line, -1)

		textData := lineData[1:]
		num := ""
		for _, text := range textData {
			num += text
		}
		data, _ := strconv.Atoi(num)

		if lineData[0] == "Time:" {
			times = []int{data}
		} else if lineData[0] == "Distance:" {
			distances = []int{data}
		}

		adder(len(line))
	}

	fmt.Printf("%#v\n", times)
	fmt.Printf("%#v\n", distances)

	for i := range times {
		cf := make(chan int)
		cl := make(chan int)
		go findFirstStrategy(times[i], distances[i], cf)
		go findLastStrategy(times[i], distances[i], cl)
		f, l := <-cf, <-cl
		possibilities := l - f + 1
		multiplier(possibilities)
	}

	fmt.Printf("Sum: %d\n", adder(0))
	fmt.Printf("Product: %d\n", multiplier(1))
}
