package day7

import (
	"AoC2025/utils"
	"math/big"
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
	expected := big.NewInt(40)
	result := SolveSecondPart(lines)
	if result.Cmp(expected) != 0 {
		t.Errorf("SolveSecondPart() = %v; want %v", result, expected)
	}
}
