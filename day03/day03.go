package day03

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type passwordLine struct {
}

// parseInput translates password input file into []passwordLine
func parseInput(filename string) ([][]string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return [][]string{}, err
	}
	input := strings.Split(string(body), "\n")
	stringGrid := make([][]string, len(input))

	for i, val := range input {
		stringGrid[i] = strings.Split(val, "")
	}
	return stringGrid, nil
}

// Part1 returns # of valid passwords
func Part1(filename string) (int, error) {
	stringGrid, err := parseInput(filename)
	y := 3
	count := 0
	for x := 1; x < len(stringGrid); x++ {
		if y >= len(stringGrid[0]) {
			y = y - len(stringGrid[0])
		}
		if stringGrid[x][y] == "#" {
			stringGrid[x][y] = "X"
			count++
		} else {
			stringGrid[x][y] = "O"
		}
		y += 3
	}

	fmt.Println("count:", count)

	return count, err
}

// // Part2 returns # of valid passwords based on positions rule
// func Part2(filename string) (int, error) {
// 	passwordLine, err := parseInput(filename)
// 	if err != nil {
// 		return 0, err
// 	}
// 	count := 0
// 	for _, line := range passwordLine {
// 		if (string(line.password[line.low-1]) == line.letter) != (string(line.password[line.high-1]) == line.letter) {
// 			count++
// 		}
// 	}
// 	return count, nil
// }
