package day05

import (
	"fmt"
	"io/ioutil"
)

// Part1 ..
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	fmt.Println(body)

	return 0, nil
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
