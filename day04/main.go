package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "-")
	inputNumber := make([]int, 0, len(input))

	for _, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		inputNumber = append(inputNumber, n)
	}

	count := 0

	// value within range
	for i := inputNumber[0]; i <= inputNumber[1]; i++ {
		str := strconv.Itoa(i)
		var prev rune
		decrease := true
		double := false
		for pos, char := range str {
			if pos > 0 {
				if char < prev {
					decrease = false
				}
				if char == prev {
					double = true
				}
			}
			prev = char
		}
		// six-digit
		if len(str) != 6 {
			continue
		}
		// never decrease
		if !decrease {
			continue
		}
		// two adjacent digits are the same
		if !double {
			continue
		}
		count++
	}

	fmt.Println("part one:", count) // 1767
}
