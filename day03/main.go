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

var center = point{0, 0}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n")
	minDistance := math.MaxInt16
	minSteps := math.MaxInt16

	wire1 := calculateWire(input[0])
	wire2 := calculateWire(input[1])

	for point, steps1 := range wire1 {
		if steps2, ok := wire2[point]; ok {
			distance := calculateDistance(center, point)
			steps := steps1 + steps2
			if distance < minDistance {
				minDistance = distance
			}
			if steps < minSteps {
				minSteps = steps
			}
		}
	}

	fmt.Println("part one:", minDistance) // 403
	fmt.Println("part two:", minSteps)    // 4158
}

func getNextPoint(direction string, lastPoint point) point {
	switch direction {
	case "R":
		return point{lastPoint.x + 1, lastPoint.y}
	case "U":
		return point{lastPoint.x, lastPoint.y + 1}
	case "L":
		return point{lastPoint.x - 1, lastPoint.y}
	case "D":
		return point{lastPoint.x, lastPoint.y - 1}
	default:
		return lastPoint
	}
}

func calculateWire(path string) map[point]int {
	wire := make(map[point]int)
	lastPoint := center
	instructions := strings.Split(path, ",")
	step := 1
	for _, command := range instructions {
		direction := command[:1]
		steps, _ := strconv.Atoi(command[1:])
		i := 1
		for i <= steps {
			lastPoint = getNextPoint(direction, lastPoint)
			if _, ok := wire[lastPoint]; !ok {
				wire[lastPoint] = step
			}
			i++
			step++
		}
	}
	return wire
}

func calculateDistance(from point, to point) int {
	return int(math.Abs(float64(to.x-from.x)) + math.Abs(float64(to.y-from.y)))
}
