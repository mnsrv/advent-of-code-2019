package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type point struct {
	x int
	y int
}
type pointTo struct {
	x        int
	y        int
	distance float64
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
	var station point

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
			station = coord1
		}
	}

	fmt.Println("part one:", max) // 256

	angles := make(map[float64][]pointTo)
	for coord := range all {
		if coord == station {
			continue
		}
		angle := getAngle(coord, station)
		distance := getDistance(coord, station)

		angles[angle] = append(angles[angle], pointTo{coord.x, coord.y, distance})
		sort.Slice(angles[angle], func(i, j int) bool { return angles[angle][i].distance < angles[angle][j].distance })
	}

	degrees := make([]float64, len(angles))
	i := 0
	for angle := range angles {
		degrees[i] = angle
		i++
	}
	sort.Float64s(degrees)

	counter := 0
	var part2 pointTo
	for _, degree := range degrees {
		if len(angles[degree]) == 0 {
			continue
		}
		counter++

		var el pointTo
		el, angles[degree] = angles[degree][0], angles[degree][1:]
		if counter == 200 {
			part2 = el
		}
	}

	fmt.Println("part two:", part2.x*100+part2.y) // 1707
}

func toPrecision(float float64, base int) float64 {
	precision := math.Pow10(base)
	return math.Floor(float*precision) / precision
}

func getAngle(point1 point, point2 point) float64 {
	deltaX := float64(point2.x) - float64(point1.x)
	deltaY := float64(point2.y) - float64(point1.y)
	angle := math.Atan2(deltaY, deltaX)*180/math.Pi - 90
	if angle < 0 {
		angle += 360
	}

	return toPrecision(angle, 2) // in degrees .00
}
func getDistance(point1 point, point2 point) float64 {
	deltaX := float64(point2.x) - float64(point1.x)
	deltaY := float64(point2.y) - float64(point1.y)

	return math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
}
