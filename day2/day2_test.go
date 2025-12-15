package day2

import (
	"AoC2025/utils"
	"testing"
)

func TestExampleInputFirstPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveFirstPart(lines[0])

	if result != 1227775554 {
		t.Fatalf("Example Input gave incorrect result it is %d, should be 1227775554", result)
	}
}

func TestExampleInputSecondPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}
	result := SolveSecondPart(lines[0])

	if result != 4174379265 {
		t.Fatalf("Example Input gave incorrect result it is %d, should be 4174379265", result)
	}
}
