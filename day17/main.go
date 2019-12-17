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
	fmt.Println("part one:", part1()) // 5724
}

func part1() int {
	input, _ := ioutil.ReadFile("input")
	inputStrings := strings.Split(string(input), ",")
	data := make([]int, 0, len(inputStrings))
	for _, s := range inputStrings {
		n, _ := strconv.Atoi(s)
		data = append(data, n)
	}

	channel := make(chan int)
	area := make(map[point]string)
	x := 0
	y := 0
	go computer(data, channel)
	for {
		output := <-channel
		str := string(output)
		area[point{x, y}] = str
		fmt.Print(str)

		x++
		if output == 10 {
			x = 0
			y++
		}

		if output == 9999 {
			break
		}
	}
	fmt.Println()

	answer := 0
	for k, v := range area {
		if v != "#" {
			continue
		}

		top := area[point{k.x, k.y - 1}]
		bottom := area[point{k.x, k.y + 1}]
		left := area[point{k.x - 1, k.y}]
		right := area[point{k.x + 1, k.y}]

		if top == "#" && bottom == "#" && left == "#" && right == "#" {
			answer += k.x * k.y
		}
	}

	return answer
}

func computer(array []int, channel chan int) {
	index := 0
	commands := make([]int, math.MaxInt16)
	copy(commands, array)
	relativeBase := 0

	getIndex := func(input []int, index int, mode int, relativeBase int) int {
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
			commands[getIndex(commands, index+1, mode1, relativeBase)] = <-channel
			// fmt.Println("INPUT", commands[getIndex(commands, index+1, mode1, relativeBase)])
			index += 2
		case 4:
			mode1 := params % 10
			channel <- commands[getIndex(commands, index+1, mode1, relativeBase)]
			// fmt.Println("OUTPUT", commands[getIndex(commands, index+1, mode1, relativeBase)])
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
			// fmt.Println("HALT")
			channel <- 9999
			break Loop
		default:
			fmt.Println("WRONG", opcode)
			break Loop
		}
	}
}
