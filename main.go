package main

import (
	"AoC2025/day1"
	"AoC2025/utils"
	"fmt"
	"os"
)

func solveDay1() {
	lines, err := utils.LoadFileAsLines("day1\\input.txt")

	if err != nil {
		panic(err)
	}

	result := day1.Solve(lines)

	fmt.Println(result)
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
		}
	}
}
