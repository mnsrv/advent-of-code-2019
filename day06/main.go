package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputStrings := strings.Split(string(data), "\n")

	orbits := make(map[string]string)
	for _, value := range inputStrings {
		relations := strings.Split(value, ")")
		parent := relations[0]
		node := relations[1]

		orbits[node] = parent
	}

	all := make([][]string, 0, len(orbits))
	for node := range orbits {
		all = append(all, getTree(orbits, node, []string{}))
	}

	total := 0
	for _, path := range all {
		total += len(path)
	}

	fmt.Println("part one:", total) // 312697

	you := getTree(orbits, "YOU", []string{})
	santa := getTree(orbits, "SAN", []string{})
	similar := 0
	i := 1
	for i <= len(you) {
		if you[len(you)-i] != santa[len(santa)-i] {
			break
		}
		similar++
		i = i + 1
	}
	santaUnique := len(santa) - 1 - similar
	youUnique := len(you) - 1 - similar
	fmt.Println("part two:", santaUnique+youUnique) // 440 is too low
}

func getTree(m map[string]string, node string, acc []string) []string {
	if _, ok := m[node]; !ok {
		return acc
	}
	return getTree(m, m[node], append(acc, node))
}
