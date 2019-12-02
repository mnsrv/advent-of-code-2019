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
	sum := 0

	for _, str := range input {
		sum += findFuel(str)
	}

	fmt.Println("sum:", sum)
}

func findFuel(str string) int {
	mass, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return int(math.Floor(float64(mass)/3)) - 2
}
