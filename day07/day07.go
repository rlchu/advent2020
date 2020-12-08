package day07

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Part1 ...
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	groups := strings.Split(string(body), "\n\n")
	fmt.Println(groups)

	return 0, nil
}

// Part2 ...
// func Part2(filename string) (int, error) {
// 	body, _ := ioutil.ReadFile(filename)
// 	groups := strings.Split(string(body), "\n\n")

// 	return answer, nil
// }
