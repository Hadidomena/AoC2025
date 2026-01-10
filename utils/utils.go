package utils

import (
	"os"
	"reflect"
	"strings"
)

// Takes in path to the file, returns all of the lines loaded from it as an []string.
func LoadFileAsLines(filename string) ([]string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(fileContent), "\n"), nil
}

// Returns absolute value of integer.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Generic function for finding if element occures in a structure.
// Supported structures are slice, array and map.
func IsIn(needle interface{}, haystack interface{}) bool {
	if needle == nil {
		return false
	}
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == needle {
				return true
			}
		}
	case reflect.Map:
		needleVal := reflect.ValueOf(needle)
		if needleVal.Type().AssignableTo(val.Type().Key()) {
			if val.MapIndex(needleVal).IsValid() {
				return true
			}
		}
	}
	return false
}

func Min[T int | int64 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int64 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MinSlice(slice []int64) int64 {
	if len(slice) == 0 {
		panic("minSlice called with empty slice")
	}
	minVal := slice[0]
	for _, val := range slice {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func MaxSlice(slice []int64) int64 {
	if len(slice) == 0 {
		panic("maxSlice called with empty slice")
	}
	maxVal := slice[0]
	for _, val := range slice {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
