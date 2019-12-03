package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := strings.Split(string(data), ",")
	input := make([]int, 0, len(inputStrings))

	for _, s := range inputStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	noun := 0
	verb := 0

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			output := computer(input, i, j)
			if output == 19690720 {
				noun = i
				verb = j
				break
			}
		}
	}

	fmt.Println("part one:", computer(input, 12, 2))
	fmt.Println("part two:", 100*noun+verb)
}

func computer(array []int, noun int, verb int) int {
	index := 0
	input := make([]int, len(array))
	copy(input, array)

	input[1] = noun
	input[2] = verb

Loop:
	for {
		// fmt.Println("step", index, input)
		opcode := input[index]
		result := 0

		switch opcode {
		case 1:
			result = input[input[index+1]] + input[input[index+2]]
		case 2:
			result = input[input[index+1]] * input[input[index+2]]
		case 99:
			// fmt.Println("HALT", input[0])
			break Loop
		default:
			// fmt.Println("WRONG")
			break Loop
		}
		input[input[index+3]] = result
		// fmt.Println("result", result, input)

		index += 4
	}

	return input[0]
}
