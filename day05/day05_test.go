package day05

import (
	"testing"
)

type testCase struct {
	fileName string
	want     int
}

var testData01 = []testCase{
	{"day05_example.txt", 820},
	{"day05_input.txt", 947},
}

var testData02 = []testCase{
	// {"day05_example.txt", 2},
	// {"day05_input.txt", 175},
}

func TestPart1(t *testing.T) {
	for _, val := range testData01 {
		answer, err := Part1(val.fileName)
		if err != nil {
			t.Fatalf("%s", err)
		}
		if answer != val.want {
			t.Fatalf("wanted: %v got: %v", val.want, answer)
		}
	}
}

// func TestPart2(t *testing.T) {
// 	for _, val := range testData02 {
// 		answer, err := Part2(val.fileName)
// 		if err != nil {
// 			t.Fatalf("%s", err)
// 		}
// 		if answer != val.want {
// 			t.Fatalf("wanted: %v got: %v", val.want, answer)
// 		}
// 	}
// }
