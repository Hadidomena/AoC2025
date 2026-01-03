package day4

func withinBounds(x, y, boundX, boundY int) bool {
	return x >= 0 && x < boundX && y >= 0 && y < boundY
}
func SolveFirstPart(lines []string) int {
	result := 0
	boundX := len(lines)
	boundY := len(lines[0])
	for i := 0; i < boundX; i++ {
		for j := 0; j < boundY; j++ {
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

func SolveSecondPart(lines []string) int64 {
	return 0
}
