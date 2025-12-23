package day3

import (
	"AoC2025/utils"
	"testing"
)

func TestExampleInputFirstPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveFirstPart(lines)
	if result != 357 {
		t.Fatalf("Example Input gave incorrect result it is %d, should be 357", result)
	}
}

func TestExampleInputSecondPart(t *testing.T) {
}
