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
	index := 0

	for _, s := range inputStrings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	input[1] = 12
	input[2] = 2

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
			fmt.Println("HALT", input[0])
			break Loop
		default:
			fmt.Println("WRONG")
			break Loop
		}
		input[input[index+3]] = result
		// fmt.Println("result", result, input)

		index += 4
	}
}
