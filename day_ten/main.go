package main

import (
	"os"
	"strings"
)

type Move int
type Direction int

const (
	COUNTER_CLOCKWISE Direction = iota
	CLOCKWISE
)

const (
	UP    Move = -141
	DOWN  Move = 141
	LEFT  Move = -1
	RIGHT Move = 1
)

type Pointer struct {
	previous Move
	input    string
	current  int
	moves    int
}

func (p *Pointer) move() {
	move := getNextMove(p.previous, rune(p.input[p.current]))
	p.previous = Move(p.current - (p.current + move))
	p.current = p.current + move
	p.moves++
}

func findStartingPoints(start int, input string) (int, int) {
	up := start + int(UP)
	left := start + int(LEFT)
	right := start + int(RIGHT)
	down := start + int(DOWN)

	println(string(input[up]))

	switch input[up] {
	case '|':
		fallthrough
	case 'F':
		fallthrough
	case '7':
		{
			switch input[left] {
			case '-':
				fallthrough
			case 'F':
				fallthrough
			case 'L':
				{
					return up, left
				}
			}
			switch input[right] {
			case '7':
				fallthrough
			case 'J':
				fallthrough
			case '.':
				{
					return up, right
				}
			}
			switch input[down] {
			case '|':
				fallthrough
			case 'J':
				fallthrough
			case 'L':
				{
					return up, down
				}

			}
		}
	}
	println(string(input[left]))
	switch input[left] {
	case '-':
		fallthrough
	case 'L':
		fallthrough
	case 'F':
		{
			println("should match")
			switch input[right] {
			case '7':
				fallthrough
			case 'J':
				fallthrough
			case '-':
				{
					return left, right
				}
			}
			switch input[down] {
			case '|':
				fallthrough
			case 'J':
				fallthrough
			case 'L':
				{
					return left, down
				}
			}
		}
	}

	switch input[right] {
	case '-':
		fallthrough
	case '7':
		fallthrough
	case 'J':
		{
			switch input[down] {
			case '|':
				fallthrough
			case 'L':
				fallthrough
			case 'J':
				{
					return right, down
				}
			}
		}
	}

	return 0, 0
}

func getNextMove(previous Move, ch rune) int {
	switch ch {
	case '-':
		{
			if previous == RIGHT {
				return int(LEFT)
			}
			return int(RIGHT)
		}
	case '|':
		{
			if previous == DOWN {
				return int(UP)
			}
			return int(DOWN)
		}
	case 'F':
		{
			if previous == DOWN {
				return int(RIGHT)
			}
			return int(DOWN)

		}
	case 'L':
		{
			if previous == RIGHT {
				return int(UP)
			}
			return int(RIGHT)

		}
	case '7':
		{
			if previous == DOWN {
				return int(LEFT)
			}
			return int(DOWN)

		}
	case 'J':
		{
			if previous == UP {
				return int(LEFT)
			}
			return int(UP)
		}
	}
	return 0
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Can't read file: ", os.Args[1])
		os.Exit(1)
	}

	input := string(b)
	start := strings.Index(input, "S")
	p1, p2 := findStartingPoints(start, input)

	ptr1 := Pointer{previous: Move(start - p1), current: p1, input: input, moves: 1}
	ptr2 := Pointer{previous: Move(start - p2), current: p2, input: input, moves: 1}
	for ptr1.current != ptr2.current {
		ptr1.move()
		if ptr1.current == ptr2.current {
			break
		}
		ptr2.move()
	}
	println("Result: ", ptr1.moves)
}
