package main

import (
	"AoC2025/day1"
	"AoC2025/day2"
	"AoC2025/day3"
	"AoC2025/day4"
	"AoC2025/utils"
	"fmt"
	"os"
)

type daySolver struct {
	solveFirstPart  func([]string) int64
	solveSecondPart func([]string) int64
}

var solvers = map[string]daySolver{
	"day1": {
		solveFirstPart: func(lines []string) int64 {
			return int64(day1.SolveFirstPart(lines))
		},
		solveSecondPart: func(lines []string) int64 {
			return int64(day1.SolveSecondPart(lines))
		},
	},
	"day2": {
		solveFirstPart: func(lines []string) int64 {
			return day2.SolveFirstPart(lines[0])
		},
		solveSecondPart: func(lines []string) int64 {
			return day2.SolveSecondPart(lines[0])
		},
	},
	"day3": {
		solveFirstPart: func(lines []string) int64 {
			return int64(day3.SolveFirstPart(lines))
		},
		solveSecondPart: day3.SolveSecondPart,
	},
	"day4": {
		solveFirstPart: func(lines []string) int64 {
			return int64(day4.SolveFirstPart(lines))
		},
		solveSecondPart: day4.SolveSecondPart,
	},
}

func solveDay(day string) {
	lines, err := utils.LoadFileAsLines(day + "\\input.txt")
	if err != nil {
		panic(err)
	}

	solver, ok := solvers[day]
	if !ok {
		fmt.Printf("Solver for day %s not found\n", day)
		return
	}

	result1 := solver.solveFirstPart(lines)
	result2 := solver.solveSecondPart(lines)

	fmt.Printf("Result for first part is %d and result for second part is %d\n", result1, result2)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <list of days you want to execute>")
		os.Exit(1)
	}

	for _, day := range args {
		if utils.IsIn(day, solvers) {
			solveDay(day)
		} else {
			fmt.Printf("Solver for day %s not found\n", day)
		}
	}
}
