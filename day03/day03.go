package day03

import (
	"io/ioutil"
	"strings"
)

type stringGrid [][]string

// parseInput translates slope input file into stringGrid
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

func (sg stringGrid) countTrees(right, down int) int {
	height := len(sg)
	width := len(sg[0])
	y := down
	count := 0

	for x := right; x < height; x += right {
		if y >= width {
			y = y - width
		}
		if sg[x][y] == "#" {
			count++
		}
		y += down
	}
	return count
}

// Part1 returns # of trees encountered across a single slope (1,3)
func Part1(filename string) (int, error) {
	stringGrid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	count := stringGrid.countTrees(1, 3)
	return count, nil
}

// Part2 returns product trees encountered across multiple slopes
func Part2(filename string) (int, error) {
	stringGrid, err := parseInput(filename)
	if err != nil {
		return 0, err
	}
	count := 1
	count *= stringGrid.countTrees(1, 1)
	count *= stringGrid.countTrees(1, 3)
	count *= stringGrid.countTrees(1, 5)
	count *= stringGrid.countTrees(1, 7)
	count *= stringGrid.countTrees(2, 1)

	return count, nil
}
