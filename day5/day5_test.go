package day5

import (
	"AoC2025/utils"
	"testing"
)

func TestFirstPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveFirstPart(lines)
	if result != 3 {
		t.Errorf("Example Input gave incorrect result it is %d, should be %d", result, 3)
	}
}

func TestSecondPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveSecondPart(lines)
	if result != 3 {
		t.Errorf("Example Input gave incorrect result it is %d, should be %d", result, 3)
	}
}
