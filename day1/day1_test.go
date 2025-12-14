package day1

import (
	"AoC2025/utils"
	"testing"
)

func TestExampleInput(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := Solve(lines)

	if result != 3 {
		t.Fatalf("Expected 3, got %d", result)
	}
}
