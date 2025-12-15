package utils

import (
	"os"
	"strings"
	"testing"
)

func TestAbsInt(t *testing.T) {
	negative, positive := -1, 1
	if AbsInt(negative) != 1 {
		t.Fatalf("Expected 1, got %d", AbsInt(negative))
	}

	if AbsInt(positive) != 1 {
		t.Fatalf("Expected 1, got %d", AbsInt(positive))
	}
}

func TestReadingLines(t *testing.T) {
	testText := "first line\nsecond line\nthird line"
	os.WriteFile("tmp.txt", []byte(testText), 0644)
	defer os.Remove("tmp.txt")

	readText, err := LoadFileAsLines("tmp.txt")

	if err != nil {
		t.Fatalf("Error loading file: %s", err)
	}

	if len(readText) != 3 {
		t.Fatalf("Expected 3 lines, got %d", len(readText))
	}
	for x := 0; x < len(readText); x++ {
		if readText[x] != strings.Split(testText, "\n")[x] {
			t.Fatalf("Expected %s, got %s", strings.Split(testText, "\n")[x], readText[x])
		}
	}
}
