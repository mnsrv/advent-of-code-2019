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
}

func getTree(m map[string]string, node string, acc []string) []string {
	if _, ok := m[node]; !ok {
		return acc
	}
	return getTree(m, m[node], append(acc, node))
}
