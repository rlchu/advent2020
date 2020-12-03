package day03

import (
	"io/ioutil"
	"strings"
)

type passwordLine struct {
}

// parseInput translates password input file into []passwordLine
func parseInput(filename string) ([]string, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	input := strings.Split(string(body), "\n")
	return input, nil
}

// Part1 returns # of valid passwords
func Part1(filename string) (int, error) {
	_, err := parseInput(filename)

	return 0, err
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
