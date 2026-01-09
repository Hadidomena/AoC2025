package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveFirstPart(lines []string) int64 {
	result := int64(0)
	splitLines := [][]string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		filtered := []string{}
		for _, part := range parts {
			if part != "" {
				filtered = append(filtered, part)
			}
		}
		splitLines = append(splitLines, filtered)
	}
	numOfLines := len(splitLines)
	lenOfLines := len(splitLines[0])
	for x := range lenOfLines {
		lineVal := int64(0)
		for y := numOfLines - 2; y >= 0; y-- {
			switch splitLines[numOfLines-1][x] {
			case "*":
				if lineVal == 0 {
					lineVal = 1
				}
				a, _ := strconv.ParseInt(splitLines[y][x], 10, 64)
				lineVal *= a
			case "+":
				a, _ := strconv.ParseInt(splitLines[y][x], 10, 64)
				lineVal += a
			}
		}
		fmt.Println("Value for column ", x, " is ", lineVal)
		result += lineVal
	}
	fmt.Println(splitLines)
	return result
}

func SolveSecondPart(lines []string) int64 {
	return 0
}
