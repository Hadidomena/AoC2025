package day2

import (
	"strconv"
	"strings"
)

func findInvalid(start, end int64) int64 {
	var result int64 = 0
	var halfStart int64 = 1
	var p10 int64 = 10
	for {
		multiplier := p10 + 1
		halfEnd := p10 - 1
		if halfStart > end/multiplier {
			break
		}
		sMin := (start + multiplier - 1) / multiplier
		sMax := end / multiplier
		actualStart := max(halfStart, sMin)
		actualEnd := min(halfEnd, sMax)

		if actualStart <= actualEnd {
			count := actualEnd - actualStart + 1
			sumS := (actualStart + actualEnd) * count / 2
			result += sumS * multiplier
		}
		if end/10 < halfStart {
			break
		}
		halfStart *= 10
		p10 *= 10
	}
	return result
}
func SolveFirstPart(line string) int64 {
	var result int64 = 0
	ranges := strings.Split(line, ",")
	for _, span := range ranges {
		splitSpan := strings.Split(span, "-")

		start, err := strconv.ParseInt(splitSpan[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseInt(splitSpan[1], 10, 64)
		if err != nil {
			panic(err)
		}
		result += findInvalid(start, end)
	}
	return result
}
