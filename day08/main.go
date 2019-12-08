package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := string(data)
	input := make([]int, 0, len(inputStrings))
	for _, s := range inputStrings {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}
		input = append(input, n)
	}

	// 25 x 6
	size := 25 * 6
	layers := make([][]int, len(input)/size)
	layerIndex := 0
	layerInfo := make(map[int]map[int]int)
	count0 := 0
	count1 := 0
	count2 := 0
	min0 := size
	minIndex := 0

	for index, value := range input {
		layers[layerIndex] = append(layers[layerIndex], value)
		switch value {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		}
		if index%size == size-1 {
			layerInfo[layerIndex] = map[int]int{0: count0, 1: count1, 2: count2}
			if count0 < min0 {
				min0 = count0
				minIndex = layerIndex
			}

			count0 = 0
			count1 = 0
			count2 = 0
			layerIndex++
		}
	}

	fmt.Println("part one:", layerInfo[minIndex][1]*layerInfo[minIndex][2])
}
