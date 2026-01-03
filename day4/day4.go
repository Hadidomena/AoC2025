package day4

func withinBounds(x, y, boundX, boundY int) bool {
	return x >= 0 && x < boundX && y >= 0 && y < boundY
}
func SolveFirstPart(lines []string) int {
	result := 0
	boundX := len(lines)
	boundY := len(lines[0])
	for i := range boundX {
		for j := range boundY {
			if lines[i][j] == '@' {
				adjacent := 0
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						if withinBounds(i+x, j+y, boundX, boundY) && lines[i+x][j+y] == '@' {
							adjacent++
						}
					}
				}
				if adjacent < 4 {
					result++
				}
			}
		}
	}
	return result
}

func solveStep(lines []string) (int64, []string) {
	backupLines := make([][]rune, len(lines))
	for i := range lines {
		backupLines[i] = []rune(lines[i])
	}
	result := int64(0)
	boundX := len(lines)
	boundY := len(lines[0])
	for i := range boundX {
		for j := range boundY {
			if lines[i][j] == '@' {
				adjacent := 0
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						if withinBounds(i+x, j+y, boundX, boundY) && lines[i+x][j+y] == '@' {
							adjacent++
						}
					}
				}
				if adjacent < 4 {
					result++
					backupLines[i][j] = '.'
				}
			}
		}
	}
	result_lines := make([]string, len(backupLines))
	for i := range backupLines {
		result_lines[i] = string(backupLines[i])
	}
	return result, result_lines
}

func SolveSecondPart(lines []string) int64 {
	result, nextStep := solveStep(lines)
	change := result
	for change != 0 {
		change, nextStep = solveStep(nextStep)
		result += change
	}
	return result
}
