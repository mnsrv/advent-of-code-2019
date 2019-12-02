package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n")
	part1 := 0
	part2 := 0

	for _, str := range input {
		mass, _ := strconv.Atoi(str)
		part1 += findFuel(mass)
		part2 += findFuelForFuel(mass, 0)
	}

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func findFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func findFuelForFuel(mass int, sum int) int {
	fuel := findFuel(mass)
	if fuel > 0 {
		return findFuelForFuel(fuel, sum+fuel)
	}
	return sum
}
