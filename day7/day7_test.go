package day7

import (
	"AoC2025/utils"
	"testing"
)

func TestFirstPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")
	if err != nil {
		t.Fatalf("Failed to load example file: %v", err)
	}
	expected := int64(21)
	result := SolveFirstPart(lines)
	if result != expected {
		t.Errorf("SolveFirstPart() = %v; want %v", result, expected)
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
