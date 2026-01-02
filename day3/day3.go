package day3

func SolveFirstPart(lines []string) int {
	result := 0
	for _, line := range lines {
		lineVal, first, second := 0, -1, -1
		for x := len(line) - 1; x >= 0; x-- {
			charVal := int(line[x] - '0')
			if second == -1 {
				second = charVal
			} else if first == -1 {
				first = charVal
			} else {
				if charVal >= first {
					if first > second {
						second = first
					}
					first = charVal
				}
			}
		}
		lineVal = first*10 + second
		result += lineVal
	}
	return result
}

func SolveSecondPart(lines []string) int64 {
	result := int64(0)
	for _, line := range lines {
		selected := make([]int, 0, 12)
		startIdx := 0

		for len(selected) < 12 {
			remaining := 12 - len(selected)
			searchEnd := len(line) - remaining
			maxVal := -1
			maxIdx := -1
			for i := startIdx; i <= searchEnd; i++ {
				charVal := int(line[i] - '0')
				if charVal > maxVal {
					maxVal = charVal
					maxIdx = i
				}
			}

			if maxIdx != -1 {
				selected = append(selected, maxVal)
				startIdx = maxIdx + 1
			}
		}
		lineVal := int64(0)
		for _, v := range selected {
			lineVal = lineVal*10 + int64(v)
		}
		result += lineVal
	}
	return result
}
