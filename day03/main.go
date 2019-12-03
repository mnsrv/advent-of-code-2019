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

	wire1 := calculateWire(input[0])
	wire2 := calculateWire(input[1])

	for key := range wire1 {
		if _, ok := wire2[key]; ok {
			distance := calculateDistance(center, key)
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	fmt.Println("part one:", minDistance) // 403
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
	for _, command := range instructions {
		direction := command[:1]
		steps, _ := strconv.Atoi(command[1:])
		i := 1
		for i <= steps {
			lastPoint = getNextPoint(direction, lastPoint)
			wire[lastPoint] = 0
			i++
		}
	}
	return wire
}

func calculateDistance(from point, to point) int {
	return int(math.Abs(float64(to.x-from.x)) + math.Abs(float64(to.y-from.y)))
}
