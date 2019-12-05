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

	computer(input) // part one: 9961446
}

func getIndex(input []int, index int, mode int) int {
	// 0 == position mode
	// 1 == immediate mode
	if mode == 1 {
		return index
	}
	return input[index]
}

func computer(array []int) {
	index := 0
	input := make([]int, len(array))
	copy(input, array)

Loop:
	for {
		// fmt.Println("step", index)
		opcode := input[index]
		result := 0
		params := 0

		// fmt.Println("opcode:", opcode)
		if opcode > 99 {
			params = opcode / 100
			opcode = opcode % 100
		}
		// fmt.Println("opcode:", opcode, "params:", params)

		switch opcode {
		case 1:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			input[input[index+3]] = input[getIndex(input, index+1, mode1)] + input[getIndex(input, index+2, mode2)]
			index += 4
		case 2:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			input[input[index+3]] = input[getIndex(input, index+1, mode1)] * input[getIndex(input, index+2, mode2)]
			index += 4
		case 3:
			fmt.Printf("Enter ID of the system: ")
			fmt.Scanln(&result)
			input[input[index+1]] = result
			index += 2
		case 4:
			mode1 := params % 10
			fmt.Println("OUTPUT:", input[getIndex(input, index+1, mode1)])
			index += 2
		case 99:
			fmt.Println("HALT")
			break Loop
		default:
			fmt.Println("WRONG", opcode)
			break Loop
		}
		// fmt.Println(input)
	}
}
