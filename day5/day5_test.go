package day5

import (
	"AoC2025/utils"
	"fmt"
	"reflect"
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
	if result != int64(14) {
		t.Errorf("Example Input gave incorrect result it is %d, should be %d", result, 3)
	}
}

func TestRangeInsert(t *testing.T) {
	ranges := [][2]int64{{1, 5}, {4, 12}, {11, 100}}
	ranges = insertIntoRanges([2]int64{4, 8}, ranges)

	if !reflect.DeepEqual(ranges, [][2]int64{{1, 100}}) {
		fmt.Println(ranges, " != ", [][2]int64{{1, 100}})
		t.Errorf("Range insert failed for already sorted ranges")
	}

	ranges = [][2]int64{{51, 100}, {7, 51}, {1, 4}}
	ranges = insertIntoRanges([2]int64{4, 8}, ranges)

	if !reflect.DeepEqual(ranges, [][2]int64{{1, 100}}) {
		fmt.Println(ranges, " != ", [][2]int64{{1, 100}})
		t.Errorf("Range insert failed for not already sorted ranges")
	}
}
