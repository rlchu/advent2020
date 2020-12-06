package day05

import (
	"io/ioutil"
	"strings"
)

func getSeatID(bpass string) int {
	r := []int{}
	for i := 0; i < 128; i++ {
		r = append(r, i)
	}
	for i := 0; i < 7; i++ {
		letter := string(bpass[i])
		if letter == "F" {
			r = r[:(len(r) / 2)]
		} else {
			r = r[(len(r) / 2):]
		}
	}
	row := r[0]

	c := []int{}
	for i := 0; i < 8; i++ {
		c = append(c, i)
	}
	for i := 7; i < 10; i++ {
		letter := string(bpass[i])
		if letter == "R" {
			c = c[(len(c) / 2):]
		} else {
			c = c[:(len(c) / 2)]
		}
	}
	column := c[0]

	return row*8 + column
}

// Part1 ..
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n")
	maxSeatID := 0
	for _, val := range input {
		seatID := getSeatID(val)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID, nil
}

// // Part2 ...
// func Part2(filename string) (int, error) {
// 	body, _ := ioutil.ReadFile(filename)
// 	input := strings.Split(string(body), "\n\n")
// 	validCount := 0
// 	for _, passport := range input {
// 		if checkForValidEntries(passport) {
// 			validCount++
// 		}
// 	}

// 	return validCount, nil
// }
