package day01

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// cleanInput reads txt file input based on name passed in and converts from []byte -> []string -> []int64
func cleanInput(filename string) []int64 {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n")
	ci := make([]int64, len(input))
	for i, val := range input {
		ci[i], _ = strconv.ParseInt(val, 10, 0)
	}
	return ci
}

// Part1 find the two entries that sum to 2020
func Part1(input string) (int64, error) {
	cleanInput := cleanInput(input)

	for i, val1 := range cleanInput {
		for _, val2 := range cleanInput[i:] {
			if val1+val2 == 2020 {
				return val1 * val2, nil
			}
		}
	}
	return 0, errors.New("could not find two numbers totalling 2020")
}

// Part2 find the two entries that sum to 2020
func Part2(input string) (int64, error) {
	cleanInput := cleanInput(input)

	for i, val1 := range cleanInput {
		for j, val2 := range cleanInput[i:] {
			for _, val3 := range cleanInput[j:] {
				if val1+val2+val3 == 2020 {
					return val1 * val2 * val3, nil
				}
			}

		}
	}
	return 0, errors.New("could not find three numbers totalling 2020")
}
