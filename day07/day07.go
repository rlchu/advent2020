package day07

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Part1 ...
func Part1(filename string) (int, error) {
	body, _ := ioutil.ReadFile(filename)
	groups := strings.Split(string(body), "\n")
	fmt.Println(groups[0])

	// test := "dull black bags contain 2 shiny brown bags, 31 plaid bronze bags, 3 wavy teal bags, 3 dull chartreuse bag"

	re1 := regexp.MustCompile(`\w+ \w+`)
	re2 := regexp.MustCompile(`\d+ \w+ \w+`)
	line := make(map[string][]string)

	subjectBag := re1.FindString(groups[0])

	for _, val := range re2.FindAllString(groups[0], -1) {
		line[subjectBag] = append(line[subjectBag], val)
	}

	fmt.Println(line)

	return 0, nil
}

// Part2 ...
// func Part2(filename string) (int, error) {
// 	body, _ := ioutil.ReadFile(filename)
// 	groups := strings.Split(string(body), "\n\n")

// 	return answer, nil
// }
