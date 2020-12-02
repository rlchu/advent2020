package day01

import (
	"advent2020/helpers"
	"errors"
)

// Part1 find the two entries that sum to 2020, return their product
func Part1(input string) (int, error) {
	testData := helpers.CleanInput(input)

	for i, val1 := range testData {
		for _, val2 := range testData[i:] {
			if val1+val2 == 2020 {
				return val1 * val2, nil
			}
		}
	}
	return 0, errors.New("could not find two numbers totaling 2020")
}

// Part2 find the three entries that sum to 2020, return their product
func Part2(input string) (int, error) {
	testData := helpers.CleanInput(input)

	for i, val1 := range testData {
		for j, val2 := range testData[i:] {
			for _, val3 := range testData[j:] {
				if val1+val2+val3 == 2020 {
					return val1 * val2 * val3, nil
				}
			}

		}
	}
	return 0, errors.New("could not find three numbers totaling 2020")
}
