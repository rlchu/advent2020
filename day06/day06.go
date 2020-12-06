package day06

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Part1 finds maximum seat number
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	input := strings.Split(string(body), "\n")
	fmt.Println(input)

	return 0, nil
}

// Part2 finds seat number with gaps before and after
// func Part2(filename string) (int, error) {
// 	body, _ := ioutil.ReadFile(filename)
// 	input := strings.Split(string(body), "\n")

// 	return yourSeat, nil
// }
