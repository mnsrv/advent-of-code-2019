package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("part one:", part1()) // 34841690
	fmt.Println("part two:", part2()) // 48776785
}

func part2() string {
	// defer timeTrack(time.Now(), "part 2")
	input, _ := ioutil.ReadFile("input")
	signal := strings.Repeat(string(input), 10000)
	offset, _ := strconv.Atoi(signal[:7])
	output := []int{}

	for _, c := range signal[offset:] {
		output = append(output, int(c-'0'))
	}
	for p := 0; p < 100; p++ {
		sum := 0
		for i := len(output) - 1; i >= 0; i-- {
			sum += output[i]
			output[i] = sum % 10
		}
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output[:8])), ""), "[]")
}

func part1() string {
	// defer timeTrack(time.Now(), "part 1")
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := strings.Split(string(data), "")
	input := make([]int, 0, len(inputStrings))
	for _, s := range inputStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	output := make([]int, 0, len(input))
	for i := 0; i < 100; i++ {
		output = phase(input)
		copy(input, output)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output[:8])), ""), "[]")
}

func phase(input []int) []int {
	newInput := make([]int, 0, len(input))

	for i := range input {
		sum := 0
		pattern := getPattern(i)
		for j, w := range input {
			sum += w * pattern[j%len(pattern)]
		}
		last := int(math.Abs(float64(sum))) % 10
		newInput = append(newInput, last)
	}

	return newInput
}

func getPattern(n int) []int {
	defaultPattern := []int{0, 1, 0, -1}
	if n == 0 {
		return append(defaultPattern[1:], defaultPattern[0])
	}

	length := len(defaultPattern) * (n + 1)
	pattern := make([]int, length)

	for i := 0; i < length; i++ {
		index := i / (n + 1) % len(defaultPattern)
		pattern[i] = defaultPattern[index]
	}
	return append(pattern[1:], pattern[0])
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
