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

func SolveSecondPart(lines []string) int {
	return 0
}
