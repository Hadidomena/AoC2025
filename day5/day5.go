package day5

import (
	"AoC2025/utils"
	"sort"
	"strconv"
	"strings"
)

func insertIntoRanges(newRange [2]int64, ranges [][2]int64) [][2]int64 {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var result [][2]int64
	i := 0

	for i < len(ranges) && ranges[i][1] < newRange[0] {
		result = append(result, ranges[i])
		i++
	}
	for i < len(ranges) && ranges[i][0] <= newRange[1] {
		newRange[0] = utils.Min(newRange[0], ranges[i][0])
		newRange[1] = utils.Max(newRange[1], ranges[i][1])
		i++
	}

	result = append(result, newRange)
	for i < len(ranges) {
		result = append(result, ranges[i])
		i++
	}

	return result
}

func SolveFirstPart(lines []string) int64 {
	countingRanges, result, ranges := true, int64(0), [][2]int64{}
	for _, line := range lines {
		if line == "" {
			countingRanges = false
			continue
		}

		if countingRanges {
			parts := strings.Split(line, "-")

			start, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}

			end, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			ranges = insertIntoRanges([2]int64{start, end}, ranges)
		} else {
			ingredient, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				continue
			}

			for _, r := range ranges {
				if ingredient >= r[0] && ingredient <= r[1] {
					result++
					break
				}
			}
		}
	}
	return result
}

func SolveSecondPart(lines []string) int64 {
	countingRanges, result, ranges := true, int64(0), [][2]int64{}
	for _, line := range lines {
		if line == "" {
			countingRanges = false
			break
		}

		if countingRanges {
			parts := strings.Split(line, "-")

			start, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}

			end, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			ranges = insertIntoRanges([2]int64{start, end}, ranges)
		}
	}
	for _, r := range ranges {
		result += r[1] - r[0] + 1
	}
	return result
}
