package main

import (
	"AoC2025/day1"
	"AoC2025/day2"
	"AoC2025/day3"
	"AoC2025/day4"
	"AoC2025/day5"
	"AoC2025/day6"
	"AoC2025/day7"
	"AoC2025/utils"
	"fmt"
	"math/big"
	"os"
)

type daySolver struct {
	solveFirstPart  func([]string) int64
	solveSecondPart interface{}
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
	"day5": {
		solveFirstPart: func(lines []string) int64 {
			return int64(day5.SolveFirstPart(lines))
		},
		solveSecondPart: day5.SolveSecondPart,
	},
	"day6": {
		solveFirstPart: func(lines []string) int64 {
			return day6.SolveFirstPart(lines)
		},
		solveSecondPart: day6.SolveSecondPart,
	},
	"day7": {
		solveFirstPart: func(lines []string) int64 {
			return day7.SolveFirstPart(lines)
		},
		solveSecondPart: day7.SolveSecondPart,
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
	var result2 interface{}
	switch f := solver.solveSecondPart.(type) {
	case func([]string) int64:
		result2 = f(lines)
	case func([]string) *big.Int:
		result2 = f(lines)
	}

	fmt.Printf("Result for first part is %d and result for second part is %v\n", result1, result2)
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
