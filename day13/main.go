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

	fmt.Println("part two:", computer(input)) // 8777
}

func draw(input map[point]int) {
	maxRows := 0
	for point := range input {
		if point.y > maxRows {
			maxRows = point.y
		}
	}
	length := maxRows + 1
	rows := make([]map[int]string, length)
	for k, v := range input {
		tile := ""
		switch v {
		case 0:
			tile = " "
		case 1:
			tile = "#"
		case 2:
			tile = "X"
		case 3:
			tile = "-"
		case 4:
			tile = "O"
		}
		if rows[k.y] == nil {
			rows[k.y] = make(map[int]string)
		}
		rows[k.y][k.x] = tile
	}
	output := make([]string, length)
	for key, row := range rows {
		out := ""
		for i := 0; i < len(row); i++ {
			value := row[i]
			out += value
		}
		output[key] = out
	}
	// for _, v := range output {
	// 	fmt.Println(v)
	// }
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
	commands[0] = 2 // free mode
	output := 0
	relativeBase := 0
	outputCounter := 0
	score := 0
	var paddle point
	var ball point
	x := 0
	y := 0
	all := make(map[point]int)

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
			draw(all)
			if paddle.x < ball.x {
				result = 1
			} else if paddle.x > ball.x {
				result = -1
			} else {
				result = 0
			}
			// fmt.Printf("Enter ID of the system: ")
			// fmt.Scanln(&result)
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
				if x == -1 && y == 0 {
					score = output
				} else {
					all[point{x, y}] = output
					if output == 3 {
						paddle = point{x, y}
					}
					if output == 4 {
						ball = point{x, y}
					}
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
			// fmt.Println("HALT")
			break Loop
		default:
			// fmt.Println("WRONG", opcode)
			break Loop
		}
	}
	return score
}
