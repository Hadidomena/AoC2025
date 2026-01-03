package day4

import (
	"AoC2025/utils"
	"testing"
)

func TestExampleInputFirstPart(t *testing.T) {
	files := []string{"example1.txt", "example2.txt", "example3.txt"}
	results := []int{13, 4, 1}
	for i, file := range files {
		lines, err := utils.LoadFileAsLines(file)

		if err != nil {
			t.Fatalf("Error loading file: %s", err)
		}

		result := SolveFirstPart(lines)
		if result != results[i] {
			t.Errorf("Example Input gave incorrect result it is %d, should be %d", result, results[i])
		}
	}
}

// func TestExampleInputSecondPart(t *testing.T) {
// 	lines, err := utils.LoadFileAsLines("example1.txt")

// 	if err != nil {
// 		t.Fatalf("Error loading file: %s", err)
// 	}

// 	result := SolveSecondPart(lines)
// 	if result != 3121910778619 {
// 		t.Fatalf("Example Input gave incorrect result it is %d, should be 3121910778619", result)
// 	}
// }
