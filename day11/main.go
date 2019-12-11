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

	fmt.Println("part one:", computer(input, 0, false)) // 2141
	fmt.Println("part two:", computer(input, 1, true))  // RPJCFZKF
	// ###  ###    ##  ##  #### #### #  # ####
	// #  # #  #    # #  # #       # # #  #
	// #  # #  #    # #    ###    #  ##   ###
	// ###  ###     # #    #     #   # #  #
	// # #  #    #  # #  # #    #    # #  #
	// #  # #     ##   ##  #    #### #  # #
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

func computer(array []int, startColor int, paint bool) int {
	index := 0
	commands := make([]int, math.MaxInt16)
	copy(commands, array)
	output := 0
	relativeBase := 0
	position := point{0, 0}
	direction := 0 // degrees up down right left

	all := map[point]int{position: startColor}
	outputCounter := 0
	maxX := 0
	maxY := 0

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
			if color, ok := all[position]; ok {
				commands[getIndex(commands, index+1, mode1, relativeBase)] = color
			} else {
				commands[getIndex(commands, index+1, mode1, relativeBase)] = 0 // black by default
			}
			index += 2
		case 4:
			mode1 := params % 10
			outputCounter++
			if outputCounter%2 == 1 {
				// new color
				all[position] = commands[getIndex(commands, index+1, mode1, relativeBase)]
			} else {
				// rotate
				// 0 left 90
				// 1 right 90
				rotate := commands[getIndex(commands, index+1, mode1, relativeBase)]
				if rotate == 0 {
					direction = direction - 90
				} else if rotate == 1 {
					direction = direction + 90
				}
				x := position.x
				y := position.y

				switch direction % 360 {
				case 0:
					y--
				case 90, -270:
					x++
				case 180, -180:
					y++
				case 270, -90:
					x--
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
				position = point{x, y}
			}
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
			if paint {
				for i := 0; i <= maxY; i++ {
					for j := 0; j <= maxX; j++ {
						if all[point{j, i}] == 1 {
							fmt.Print("#")
						} else {
							fmt.Print(" ")
						}
					}
					fmt.Print("\n")
				}
			}
			output = len(all)
			break Loop
		default:
			fmt.Println("WRONG", opcode)
			break Loop
		}
	}
	return output
}
