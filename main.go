package main

import (
	"AoC2025/day1"
	"AoC2025/day2"
	"AoC2025/day3"
	"AoC2025/day4"
	"AoC2025/day5"
	"AoC2025/day6"
	"AoC2025/day7"
	"AoC2025/day8"
	"AoC2025/day9"
	"AoC2025/utils"
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"log"
	"math/big"
	"os"
	"sort"
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
	"day8": {
		solveFirstPart: func(lines []string) int64 {
			return int64(day8.SolvePart1(lines, 1000))
		},
		solveSecondPart: func(lines []string) int64 {
			return day8.SolvePart2(lines)
		},
	},
	"day9": {
		solveFirstPart: func(lines []string) int64 {
			return day9.SolveFirstPart(lines)
		},
		solveSecondPart: day9.SolveSecondPart,
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

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	var choices []string
	for day := range solvers {
		choices = append(choices, day)
	}
	sort.Strings(choices)

	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select the days you want to execute:\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress space to select, enter to run, q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok {
		var selectedDays []string
		for i := range m.selected {
			selectedDays = append(selectedDays, m.choices[i])
		}

		if len(selectedDays) == 0 {
			fmt.Println("No days selected.")
			return
		}

		sort.Strings(selectedDays)
		for _, day := range selectedDays {
			fmt.Printf("--- Solving %s ---\n", day)
			solveDay(day)
			fmt.Println()
		}
	}
}
