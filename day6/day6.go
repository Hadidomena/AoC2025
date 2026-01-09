package day6

import (
	"fmt"
	"strings"
)

func SolveFirstPart(lines []string) int64 {
	result := int64(0)
	splitLines := [][]string{}
	for _, line := range lines {
		splitLines = append(splitLines, strings.Split(line, " "))
	}
	fmt.Println(splitLines)
	return result
}

func SolveSecondPart(lines []string) int64 {
	return 0
}
