package main

import (
	"AoC2025/day1"
	"AoC2025/day2"
	"AoC2025/day3"
	"AoC2025/utils"
	"fmt"
	"os"
)

func solveDay1() {
	lines, err := utils.LoadFileAsLines("day1\\input.txt")

	if err != nil {
		panic(err)
	}

	result1, result2 := day1.Solve(lines), day1.SolveSecondPart(lines)

	fmt.Println("Result for first part is", result1, "and result for second part is", result2)
}

func solveDay2() {
	lines, err := utils.LoadFileAsLines("day2\\input.txt")

	if err != nil {
		panic(err)
	}
	result1 := day2.SolveFirstPart(lines[0])
	result2 := day2.SolveSecondPart(lines[0])

	fmt.Printf("Result for first part is %d and result for second part is %d", result1, result2)
}

func solveDay3() {
	lines, err := utils.LoadFileAsLines("day3\\input.txt")

	if err != nil {
		panic(err)
	}
	result1 := day3.SolveFirstPart(lines)
	result2 := day3.SolveSecondPart(lines)

	fmt.Printf("Result for first part is %d and result for second part is %d", result1, result2)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go <list of days you want to execute>")
		os.Exit(1)
	}

	for x := 0; x < len(args); x++ {
		switch args[x] {
		case "day1":
			solveDay1()
		case "day2":
			solveDay2()
		case "day3":
			solveDay3()
		}
	}
}
