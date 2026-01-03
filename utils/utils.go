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
