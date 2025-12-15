package day1

import (
	"AoC2025/utils"
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

func SolveSecondPart(lines []string) int {
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
			result += (position - 1) / 100
			position %= 100
		case 'L':
			if position == 0 {
				result--
			}
			position -= val
			result += utils.AbsInt((position + 1) / 100)
			if position < 0 && position%100 == 0 {
				result++
			}
			position %= 100
		}
		if position == 0 {
			result++
		}
		if position < 0 {
			position += 100
			result++
		}
	}
	return result
}
