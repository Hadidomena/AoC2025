package day6

import (
	"AoC2025/utils"
	"testing"
)

func TestSolveFirstPart(t *testing.T) {
	testFiles, testResults := []string{"example.txt", "example_with_irregular_spaces.txt"}, []int64{4277556, 4277556}
	for i, file := range testFiles {
		lines, err := utils.LoadFileAsLines(file)
		if err != nil {
			t.Fatalf("Failed to load example file: %v", err)
		}
		expected := testResults[i]
		result := SolveFirstPart(lines)
		if result != expected {
			t.Errorf("SolveFirstPart() = %v; want %v", result, expected)
		}
	}
}

func TestSolveSecondPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")
	if err != nil {
		t.Fatalf("Failed to load example file: %v", err)
	}
	expected := int64(3263827)
	result := SolveSecondPart(lines)
	if result != expected {
		t.Errorf("SolveSecondPart() = %v; want %v", result, expected)
	}
}
