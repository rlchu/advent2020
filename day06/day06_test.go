package day06

import (
	"testing"
)

type testCase struct {
	fileName string
	want     int
}

var testData01 = []testCase{
	{"day06_example.txt", -1},
	// {"day06_input.txt", 947},
}

var testData02 = []testCase{
	// {"day06_example.txt", -1},
	// {"day06_input.txt", 636},
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
