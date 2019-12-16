package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
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

	fmt.Println("part one:", part1(input, 100))
}

func part1(input []int, count int) string {
	output := make([]int, 0, len(input))
	for i := 0; i < count; i++ {
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
