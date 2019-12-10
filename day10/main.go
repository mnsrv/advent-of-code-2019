package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	inputStrings := strings.Split(string(data), "\n")
	all := make(map[point]int)

	for y, row := range inputStrings {
		for x, el := range row {
			if string(el) == "#" {
				all[point{x, y}] = 0
			}
		}
	}

	max := 0

	for coord1 := range all {
		angles := make(map[float64]int)
		for coord2 := range all {
			if coord1 == coord2 {
				continue
			}
			angle := getAngle(coord1, coord2)
			if _, ok := angles[angle]; !ok {
				all[coord1]++
				angles[angle] = 1
			}
		}
		if all[coord1] > max {
			max = all[coord1]
		}
	}

	fmt.Println("part one:", max) // 256
}

func getAngle(point1 point, point2 point) float64 {
	deltaX := float64(point2.x) - float64(point1.x)
	deltaY := float64(point2.y) - float64(point1.y)

	return math.Atan2(deltaY, deltaX) // in radians
}
