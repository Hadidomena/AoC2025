package day6

import (
	"AoC2025/utils"
	"strconv"
	"strings"
)

func splitIntoColumns(lines []string) [][]string {
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
	return splitLines
}

func addOrMultiply(isAdd, isMul bool, values []int64) int64 {
	if isAdd {
		sum := int64(0)
		for _, val := range values {
			sum += val
		}
		return sum
	}
	if isMul {
		product := int64(1)
		for _, val := range values {
			product *= val
		}
		return product
	}
	return 0
}

func SolveFirstPart(lines []string) int64 {
	result := int64(0)
	splitLines := splitIntoColumns(lines)
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
		result += lineVal
	}
	return result
}

func SolveSecondPart(lines []string) int64 {
	result := int64(0)
	numOfLines := len(lines)
	lenOfLines := len(lines[0])
	isMul, isAdd := false, false
	values, howManySpaces := []int64{}, 0
	for x := lenOfLines - 1; x >= 0; x-- {
		colVal := []rune{}
		for y := numOfLines - 1; y >= 0; y-- {
			switch lines[y][x] {
			case '*':
				isMul = true
				howManySpaces = 0
			case '+':
				isAdd = true
				howManySpaces = 0
			case ' ':
				howManySpaces++
			default:
				colVal = append(colVal, rune(lines[y][x]))
				howManySpaces = 0
			}
		}
		if howManySpaces >= numOfLines {
			if isMul {
				mul := int64(1)
				for _, val := range values {
					mul *= val
				}
				result += mul
				isMul = false
			}
			if isAdd {
				add := int64(0)
				for _, val := range values {
					add += val
				}
				result += add
				isAdd = false
			}
			values = []int64{}
			continue
		}
		reversed := utils.StringReverse(string(colVal))
		val, _ := strconv.ParseInt(reversed, 10, 64)
		values = append(values, val)
	}
	if isMul {
		mul := int64(1)
		for _, val := range values {
			mul *= val
		}
		result += mul
		isMul = false
	}
	if isAdd {
		add := int64(0)
		for _, val := range values {
			add += val
		}
		result += add
		isAdd = false
	}
	values = []int64{}
	return result
}
