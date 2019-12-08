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
	width := 25
	height := 6
	size := width * height
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

	fmt.Println("part one:", layerInfo[minIndex][1]*layerInfo[minIndex][2]) // 1792

	image := "\n"
	for index := range layers[0] {
		pixel := getPixelColor(layers, 0, index)
		nextRow := ""
		if (index+1)%width == 0 {
			nextRow = "\n"
		}
		image = image + pixel + nextRow
	}
	fmt.Println("part two:", image) // LJECH
	//	x      xx xxxx  xx  x  x
	//	x       x x    x  x x  x
	//	x       x xxx  x    xxxx
	//	x       x x    x    x  x
	//	x    x  x x    x  x x  x
	//	xxxx  xx  xxxx  xx  x  x
}

func getPixelColor(layers [][]int, layerIndex int, index int) string {
	pixel := layers[layerIndex][index]
	switch pixel {
	case 2:
		return getPixelColor(layers, layerIndex+1, index)
	case 1:
		return "x"
	case 0:
		return " "
	default:
		return strconv.Itoa(pixel)
	}
}
