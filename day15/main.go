package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
type tree struct {
	up    *tree
	right *tree
	down  *tree
	left  *tree
	value int
}

const droid = "🤖"
const wall = "🟦"
const nowhere = ""
const dot = "⬜️"
const space = "⬛️"
const oxygen = "✅"
const startPlace = "🚩"
const stopSquare = "🟥"
const oxygenSquare = "🟩"
const preOxygenSquare = "🟨"

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

	area := make(map[point]string)
	part1(input, area)
	part1(input, area)                    // 294
	fmt.Println("part two:", part2(area)) // 388
}

func canFillWithOxygen(square string) (canFill bool) {
	// no red white start robot green
	switch square {
	case stopSquare:
		canFill = true
	case dot:
		canFill = true
	case startPlace:
		canFill = true
	case droid:
		canFill = true
	}
	return
}

func hasSpaceWithoutOxygen(area map[point]string) (hasSpace bool) {
	for _, v := range area {
		if canFillWithOxygen(v) {
			hasSpace = true
			break
		}
	}
	return
}

func part2(area map[point]string) int {
	var oxygenPoint point
	for k, v := range area {
		if v == oxygen {
			oxygenPoint = k
		}
	}

	minutes := 0
	area[oxygenPoint] = oxygenSquare
	for {
		points := make([]point, 0)
		for k, v := range area {
			if v == oxygenSquare {
				top := point{k.x + 1, k.y}
				if canFillWithOxygen(area[top]) {
					area[top] = preOxygenSquare
					points = append(points, top)
				}
				bottom := point{k.x - 1, k.y}
				if canFillWithOxygen(area[bottom]) {
					area[bottom] = preOxygenSquare
					points = append(points, bottom)
				}
				right := point{k.x, k.y + 1}
				if canFillWithOxygen(area[right]) {
					area[right] = preOxygenSquare
					points = append(points, right)
				}
				left := point{k.x, k.y - 1}
				if canFillWithOxygen(area[left]) {
					area[left] = preOxygenSquare
					points = append(points, left)
				}
			}
		}
		for _, v := range points {
			area[v] = oxygenSquare
		}
		minutes++
		printArea(area)

		hasSpace := hasSpaceWithoutOxygen(area)
		if !hasSpace {
			break
		}
	}

	return minutes
}

func part1(input []int, area map[point]string) {
	channel := make(chan int)
	start := point{0, 0}
	prevPosition := start
	position := start
	area[position] = droid
	command := 1

	go computer(input, channel)

Infinite:
	for {
		channel <- command
		output := <-channel
		nextPosition := getNextPosition(position, command, output)

		switch output {
		case 0:
			wallPosition := getNextPosition(position, command, 1)
			command = getNextCommand(command, area, prevPosition, nextPosition)
			area[wallPosition] = wall
		case 1:
			prevPosition = position
			command = getNextCommand(command, area, prevPosition, nextPosition)
			area[nextPosition] = droid
			if position == start {
				area[position] = startPlace
			} else {
				if (area[point{position.x + 1, position.y}] == wall || area[point{position.x + 1, position.y}] == droid || area[point{position.x + 1, position.y}] == stopSquare) &&
					(area[point{position.x - 1, position.y}] == wall || area[point{position.x - 1, position.y}] == droid || area[point{position.x - 1, position.y}] == stopSquare) &&
					(area[point{position.x, position.y + 1}] == wall || area[point{position.x, position.y + 1}] == droid || area[point{position.x, position.y + 1}] == stopSquare) &&
					(area[point{position.x, position.y - 1}] == wall || area[point{position.x, position.y - 1}] == droid || area[point{position.x, position.y - 1}] == stopSquare) {
					area[position] = stopSquare
				} else {
					area[position] = dot
				}
			}
		case 2:
			area[nextPosition] = oxygen
			printArea(area)
			break Infinite
		}
		printArea(area)
		position = nextPosition
	}

	close(channel)
}

func getNextCommand(command int, area map[point]string, from point, position point) (nextCommand int) {
	// 1: north ⬆
	// 2: south ⬇
	// 3: west  ⬅
	// 4: east  ⮕

	up := point{position.x, position.y + 1}
	down := point{position.x, position.y - 1}
	left := point{position.x - 1, position.y}
	right := point{position.x + 1, position.y}

	if area[up] == nowhere {
		nextCommand = 1
	} else if area[right] == nowhere {
		nextCommand = 4
	} else if area[down] == nowhere {
		nextCommand = 2
	} else if area[left] == nowhere {
		nextCommand = 3
	} else if (area[up] == dot || area[up] == startPlace || area[up] == oxygen) && from != up {
		nextCommand = 1
	} else if (area[right] == dot || area[right] == startPlace || area[right] == oxygen) && from != right {
		nextCommand = 4
	} else if (area[down] == dot || area[down] == startPlace || area[down] == oxygen) && from != down {
		nextCommand = 2
	} else if (area[left] == dot || area[left] == startPlace || area[left] == oxygen) && from != left {
		nextCommand = 3
	} else {
		fmt.Println("NEVER HERE")
		switch command {
		case 1:
			nextCommand = 4
		case 2:
			nextCommand = 3
		case 3:
			nextCommand = 1
		case 4:
			nextCommand = 2
		}
	}

	return
}

func getNextPosition(position point, command int, output int) (nextPosition point) {
	if output == 0 {
		nextPosition = position
	} else {
		x := position.x
		y := position.y

		switch command {
		case 1:
			// 1: north ⬆
			y++
		case 2:
			// 2: south ⬇
			y--
		case 3:
			// 3: west ⬅
			x--
		case 4:
			// 4: east ⮕
			x++
		}
		nextPosition = point{x, y}
	}
	return
}

func clear() {
	// clear terminal
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func printArea(area map[point]string) {
	minX := math.MaxInt16
	maxX := 0
	minY := math.MaxInt16
	maxY := 0

	for key := range area {
		if key.x < minX {
			minX = key.x
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.y > maxY {
			maxY = key.y
		}
	}

	clear()

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if val, ok := area[point{x, y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(space)
			}
		}
		fmt.Print("\n")
	}
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
			fmt.Println("HALT")
			break Loop
		default:
			fmt.Println("WRONG", opcode)
			break Loop
		}
	}
}
