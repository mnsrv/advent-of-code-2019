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

	highest := 0

	for perm := range generatePermutations([]int{0, 1, 2, 3, 4}) {
		signal := 0
		for _, phase := range perm {
			signal = computer(input, phase, signal)
		}
		if signal > highest {
			highest = signal
		}
	}

	fmt.Println("part one:", highest)
}

func generatePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func getIndex(input []int, index int, mode int) int {
	// 0 == position mode
	// 1 == immediate mode
	if mode == 1 {
		return index
	}
	return input[index]
}

func computer(array []int, phase int, input int) int {
	index := 0
	commands := make([]int, len(array))
	copy(commands, array)
	output := 0
	phaseSetted := false

Loop:
	for {
		opcode := commands[index]
		params := 0

		if opcode > 99 {
			params = opcode / 100
			opcode = opcode % 100
		}

		switch opcode {
		case 1:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			commands[commands[index+3]] = commands[getIndex(commands, index+1, mode1)] + commands[getIndex(commands, index+2, mode2)]
			index += 4
		case 2:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			commands[commands[index+3]] = commands[getIndex(commands, index+1, mode1)] * commands[getIndex(commands, index+2, mode2)]
			index += 4
		case 3:
			if !phaseSetted {
				phaseSetted = true
				commands[commands[index+1]] = phase
			} else {
				commands[commands[index+1]] = input
			}
			index += 2
		case 4:
			mode1 := params % 10
			// fmt.Println("OUTPUT:", commands[getIndex(commands, index+1, mode1)])
			output = commands[getIndex(commands, index+1, mode1)]
			index += 2
		case 5:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1)] != 0 {
				index = commands[getIndex(commands, index+2, mode2)]
			} else {
				index += 3
			}
		case 6:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1)] == 0 {
				index = commands[getIndex(commands, index+2, mode2)]
			} else {
				index += 3
			}
		case 7:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1)] < commands[getIndex(commands, index+2, mode2)] {
				commands[commands[index+3]] = 1
			} else {
				commands[commands[index+3]] = 0
			}
			index += 4
		case 8:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1)] == commands[getIndex(commands, index+2, mode2)] {
				commands[commands[index+3]] = 1
			} else {
				commands[commands[index+3]] = 0
			}
			index += 4
		case 99:
			// fmt.Println("HALT")
			break Loop
		default:
			// fmt.Println("WRONG", opcode)
			break Loop
		}
	}
	return output
}
