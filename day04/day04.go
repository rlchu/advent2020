package day04

import (
	"io/ioutil"
	"strings"
)

// parseInput ..
func parseInput(filename string) (stringGrid, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return stringGrid{}, err
	}
	input := strings.Split(string(body), "\n")
	sg := make(stringGrid, len(input))

	for i, val := range input {
		sg[i] = strings.Split(val, "")
	}
	return sg, nil
}

// Part1 ..
func Part1(filename string) (int, error) {
	stringGrid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	count := stringGrid.countTrees(1, 3)
	return count, nil
}

// Part2 ...
func Part2(filename string) (int, error) {
	stringGrid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	count := stringGrid.countTrees(1, 1)
	count *= stringGrid.countTrees(1, 3)
	count *= stringGrid.countTrees(1, 5)
	count *= stringGrid.countTrees(1, 7)
	count *= stringGrid.countTrees(2, 1)

	return count, nil
}
