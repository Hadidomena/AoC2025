package day7

import (
	"AoC2025/utils"
)

func SolveFirstPart(lines []string) int64 {
	beams, result := []int{}, int64(0)
	for i, line := range lines {
		if i == 0 {
			for j, char := range line {
				if char == 'S' {
					beams = append(beams, j)
				}
			}
		}
		newBeams := []int{}
		for _, beam := range beams {
			if line[beam] == '^' {
				result++
				if !utils.IsIn(beam-1, newBeams) && beam-1 >= 0 {
					newBeams = append(newBeams, beam-1)
				}
				if !utils.IsIn(beam+1, newBeams) && beam+1 < len(line) {
					newBeams = append(newBeams, beam+1)
				}
			} else if !utils.IsIn(beam, newBeams) {
				newBeams = append(newBeams, beam)
			}
		}
		beams = newBeams
	}
	return result
}

func SolveSecondPart(lines []string) int64 {
	return 0
}
