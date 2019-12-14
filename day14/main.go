package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type chemical struct {
	name   string
	amount int
}
type reaction struct {
	receipt []chemical
	amount  int
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := strings.Split(string(data), "\n")
	receipts := make(map[string]reaction)

	for _, input := range inputStrings {
		answer := strings.Split(input, " => ")
		right := getChemical(answer[1])
		left := strings.Split(answer[0], ", ")
		value := make([]chemical, 0, len(left))
		for _, chem := range left {
			value = append(value, getChemical(chem))
		}
		receipts[right.name] = reaction{value, right.amount}
	}

	fmt.Println("part one:", produce(receipts, 1)) // 783895
}

func produce(receipts map[string]reaction, n int) int {
	inventory := make(map[string]int)

	return triggerReaction("FUEL", n, receipts, inventory)
}

func triggerReaction(chemical string, amount int, receipts map[string]reaction, inventory map[string]int) int {
	ore := 0
	neededRatio := int(math.Ceil(float64(amount) / float64(receipts[chemical].amount)))

	for _, reactive := range receipts[chemical].receipt {
		newAmount := reactive.amount * neededRatio
		if reactive.name == "ORE" {
			ore += newAmount
		} else {
			if inventory[reactive.name] < newAmount {
				ore += triggerReaction(reactive.name, newAmount-inventory[reactive.name], receipts, inventory)
			}

			inventory[reactive.name] -= newAmount
		}
	}

	inventory[chemical] += neededRatio * receipts[chemical].amount

	return ore
}

func getChemical(str string) chemical {
	answer := strings.Split(str, " ")
	name := answer[1]
	amount, _ := strconv.Atoi(answer[0])

	return chemical{name, amount}
}
