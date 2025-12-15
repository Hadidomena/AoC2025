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

func findInvalidPartTwo(start, end int64) int64 {
	unique := make(map[int64]struct{})
	sStr := strconv.FormatInt(end, 10)
	maxLen := len(sStr)

	for L := 1; L <= maxLen/2; L++ {
		halfStart := int64(1)
		for i := 0; i < L-1; i++ {
			halfStart *= 10
		}
		halfEnd := halfStart*10 - 1

		p10L := halfStart * 10
		currentTerm := p10L
		multiplier := 1 + p10L

		for {
			if multiplier > end {
				break
			}
			if halfStart > end/multiplier {
				break
			}

			sMin := (start + multiplier - 1) / multiplier
			sMax := end / multiplier
			actualStart := max(halfStart, sMin)
			actualEnd := min(halfEnd, sMax)

			if actualStart <= actualEnd {
				for s := actualStart; s <= actualEnd; s++ {
					val := s * multiplier
					unique[val] = struct{}{}
				}
			}

			if end/p10L < currentTerm {
				break
			}
			currentTerm *= p10L
			multiplier += currentTerm
		}
	}

	var result int64
	for k := range unique {
		result += k
	}
	return result
}

func SolveSecondPart(line string) int64 {
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
		result += findInvalidPartTwo(start, end)
	}
	return result
}
