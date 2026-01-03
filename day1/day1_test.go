package day1

import (
	"AoC2025/utils"
	"strings"
	"testing"
)

func TestExampleInput(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveFirstPart(lines)

	if result != 3 {
		t.Fatalf("Expected 3, got %d", result)
	}
}

func TestSecondPart(t *testing.T) {
	lines, err := utils.LoadFileAsLines("example.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	result := SolveSecondPart(lines)

	if result != 6 {
		t.Fatalf("Expected 6, got %d", result)
	}
}

func TestLandingOnZero(t *testing.T) {
	inputRight := "R50 R100 R1000"
	inputLeft := "L50 L100 L1000"

	linesRight := strings.Split(inputRight, " ")
	linesLeft := strings.Split(inputLeft, " ")

	resultRight := SolveSecondPart(linesRight)
	resultLeft := SolveSecondPart(linesLeft)

	if resultRight != 12 {
		t.Fatalf("Expected 12, got %d", resultRight)
	}

	if resultLeft != 12 {
		t.Fatalf("Expected 12, got %d", resultLeft)
	}
}

func TestLeavingFromZero(t *testing.T) {
	inputRight := "R50 R100 R2"
	inputLeft := "L50 L100 L2"

	linesRight := strings.Split(inputRight, " ")
	linesLeft := strings.Split(inputLeft, " ")

	resultRight := SolveSecondPart(linesRight)
	resultLeft := SolveSecondPart(linesLeft)

	if resultRight != 2 {
		t.Fatalf("Expected 2, got %d", resultRight)
	}

	if resultLeft != 2 {
		t.Fatalf("Expected 2, got %d", resultLeft)
	}

}
