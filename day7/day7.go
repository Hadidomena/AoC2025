package day7

import (
	"AoC2025/utils"
	"math/big"
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

func SolveSecondPart(lines []string) *big.Int {
	beams := make(map[int]*big.Int)
	for i, line := range lines {
		if i == 0 {
			for j, char := range line {
				if char == 'S' {
					beams[j] = big.NewInt(1)
				}
			}
		} else {
			newBeams := make(map[int]*big.Int)
			for k, v := range beams {
				if line[k] == '^' {
					if k-1 >= 0 {
						if newBeams[k-1] == nil {
							newBeams[k-1] = new(big.Int)
						}
						newBeams[k-1].Add(newBeams[k-1], v)
					}
					if k+1 < len(line) {
						if newBeams[k+1] == nil {
							newBeams[k+1] = new(big.Int)
						}
						newBeams[k+1].Add(newBeams[k+1], v)
					}
				} else {
					if newBeams[k] == nil {
						newBeams[k] = new(big.Int)
					}
					newBeams[k].Add(newBeams[k], v)
				}
			}
			beams = newBeams
		}
	}
	result := new(big.Int)
	for _, v := range beams {
		result.Add(result, v)
	}
	return result
}
