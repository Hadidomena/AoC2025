package day1

import (
	"strconv"
)

func Solve(lines []string) int {
	position, result := 50, 0
	for _, line := range lines {
		dir := line[0]
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch dir {
		case 'R':
			position += val
		case 'L':
			position -= val
		}
		position %= 100
		if position == 0 {
			result++
		}
	}
	return result
}
