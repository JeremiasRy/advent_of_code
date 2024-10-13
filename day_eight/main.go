package main

import (
	"os"
	"strings"
)

type Instruction int

const (
	InstructionLeft Instruction = iota
	InstructionRight
)

type MapEntry struct {
	left  string
	right string
}

func main() {
	if len(os.Args) != 2 {
		println("usage go run main.go <input file>")
		os.Exit(2)
	}

	bytes, err := os.ReadFile(os.Args[1])

	if err != nil {
		println("Failed to open file %s", os.Args[1])
		os.Exit(1)
	}

	instructions := []Instruction{}
	m := map[string]MapEntry{}
	input := strings.Split(string(bytes), "\n")
	key := ""

	for _, ch := range input[0] {
		if ch == 'L' {
			instructions = append(instructions, InstructionLeft)
			continue
		}

		if ch == 'R' {
			instructions = append(instructions, InstructionRight)
			continue
		}
	}

	for _, str := range input[1:] {
		if len(str) == 0 {
			continue
		}
		entry := MapEntry{left: str[7:10], right: str[12:15]}
		m[str[0:3]] = entry
	}

	steps := 0
	current := m["AAA"]

	for key != "ZZZ" {
		for _, instruction := range instructions {
			steps++
			switch instruction {
			case InstructionLeft:
				key = current.left
				current = m[key]

			case InstructionRight:
				key = current.right
				current = m[key]
			}

			if key == "ZZZ" {
				break
			}
		}
	}
	println(steps)
}
