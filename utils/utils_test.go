package utils

import (
	"os"
	"strings"
	"testing"
)

func TestAbsInt(t *testing.T) {
	negative, positive, zero := -1, 1, 0
	if AbsInt(negative) != 1 {
		t.Fatalf("Expected 1, got %d", AbsInt(negative))
	}

	if AbsInt(positive) != 1 {
		t.Fatalf("Expected 1, got %d", AbsInt(positive))
	}

	if AbsInt(zero) != 0 {
		t.Fatalf("Expected 0, got %d", AbsInt(zero))
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

func TestIsIn(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	array := [5]string{"a", "b", "c", "d", "e"}
	mapping := map[string]int{"one": 1, "two": 2, "three": 3}

	if !IsIn(3, slice) {
		t.Fatalf("Expected 3 to be in slice")
	}

	if IsIn(6, slice) {
		t.Fatalf("Expected 6 to not be in slice")
	}

	if !IsIn("c", array) {
		t.Fatalf("Expected 'c' to be in array")
	}

	if IsIn("f", array) {
		t.Fatalf("Expected 'f' to not be in array")
	}

	if !IsIn("two", mapping) {
		t.Fatalf("Expected 'two' to be in map keys")
	}

	if IsIn("four", mapping) {
		t.Fatalf("Expected 'four' to not be in map keys")
	}
}

func TestMinMax(t *testing.T) {
	if Min(3, 5) != 3 {
		t.Fatalf("Expected Min(3,5) to be 3")
	}
	if Max(3, 5) != 5 {
		t.Fatalf("Expected Max(3,5) to be 5")
	}

	if Min[int64](3000000000, 5000000000) != 3000000000 {
		t.Fatalf("Expected Min(3000000000,5000000000) to be 3000000000")
	}
	if Max[int64](3000000000, 5000000000) != 5000000000 {
		t.Fatalf("Expected Max(3000000000,5000000000) to be 5000000000")
	}

	if Min(3.5, 5.2) != 3.5 {
		t.Fatalf("Expected Min(3.5,5.2) to be 3.5")
	}
	if Max(3.5, 5.2) != 5.2 {
		t.Fatalf("Expected Max(3.5,5.2) to be 5.2")
	}
}
