package day8

import (
	"AoC2025/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")
	if err != nil {
		t.Fatalf("could not load example.txt: %v", err)
	}
	want := 40
	got := SolvePart1(lines, 10)
	if got != want {
		t.Errorf("SolvePart1() = %v, want %v", got, want)
	}
}

func TestSecondPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")
	if err != nil {
		t.Fatalf("Failed to load example file: %v", err)
	}
	expected := int64(0)
	result := SolveSecondPart(lines)
	if result != expected {
		t.Errorf("SolveSecondPart() = %v; want %v", result, expected)
	}
}
