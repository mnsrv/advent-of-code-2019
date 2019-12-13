package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

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

	fmt.Println("part one:", computer(input)) // 180
}
func getIndex(input []int, index int, mode int, relativeBase int) int {
	// 0 == position mode
	// 1 == immediate mode
	// 2 == relative mode
	if mode == 2 {
		return input[index] + relativeBase
	}
	if mode == 1 {
		return index
	}
	return input[index]
}

func computer(array []int) int {
	index := 0
	commands := make([]int, math.MaxInt16)
	copy(commands, array)
	output := 0
	relativeBase := 0
	outputCounter := 0
	x := 0
	y := 0
	all := make(map[point]int)
	blockCounter := 0

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
			params = params / 10
			mode3 := params % 10
			commands[getIndex(commands, index+3, mode3, relativeBase)] = commands[getIndex(commands, index+1, mode1, relativeBase)] + commands[getIndex(commands, index+2, mode2, relativeBase)]
			index += 4
		case 2:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			params = params / 10
			mode3 := params % 10
			commands[getIndex(commands, index+3, mode3, relativeBase)] = commands[getIndex(commands, index+1, mode1, relativeBase)] * commands[getIndex(commands, index+2, mode2, relativeBase)]
			index += 4
		case 3:
			mode1 := params % 10
			result := 0
			fmt.Printf("Enter ID of the system: ")
			fmt.Scanln(&result)
			commands[getIndex(commands, index+1, mode1, relativeBase)] = result
			// fmt.Println("INPUT:", result)
			index += 2
		case 4:
			mode1 := params % 10
			output = commands[getIndex(commands, index+1, mode1, relativeBase)]
			// fmt.Println("OUTPUT:", output)
			if outputCounter%3 == 0 {
				x = output
			} else if outputCounter%3 == 1 {
				y = output
			} else {
				all[point{x, y}] = output
				if output == 2 {
					blockCounter++
				}
			}
			outputCounter++
			index += 2
		case 5:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1, relativeBase)] != 0 {
				index = commands[getIndex(commands, index+2, mode2, relativeBase)]
			} else {
				index += 3
			}
		case 6:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			if commands[getIndex(commands, index+1, mode1, relativeBase)] == 0 {
				index = commands[getIndex(commands, index+2, mode2, relativeBase)]
			} else {
				index += 3
			}
		case 7:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			params = params / 10
			mode3 := params % 10
			if commands[getIndex(commands, index+1, mode1, relativeBase)] < commands[getIndex(commands, index+2, mode2, relativeBase)] {
				commands[getIndex(commands, index+3, mode3, relativeBase)] = 1
			} else {
				commands[getIndex(commands, index+3, mode3, relativeBase)] = 0
			}
			index += 4
		case 8:
			mode1 := params % 10
			params = params / 10
			mode2 := params % 10
			params = params / 10
			mode3 := params % 10
			if commands[getIndex(commands, index+1, mode1, relativeBase)] == commands[getIndex(commands, index+2, mode2, relativeBase)] {
				commands[getIndex(commands, index+3, mode3, relativeBase)] = 1
			} else {
				commands[getIndex(commands, index+3, mode3, relativeBase)] = 0
			}
			index += 4
		case 9:
			mode1 := params % 10
			relativeBase += commands[getIndex(commands, index+1, mode1, relativeBase)]
			index += 2
		case 99:
			// fmt.Println("HALT", all, blockCounter)
			break Loop
		default:
			// fmt.Println("WRONG", opcode)
			break Loop
		}
	}
	return blockCounter
}
