package main

import (
	"os"
	"strings"
)

type SpaceNode struct {
	galaxy int
	col    int
	row    int
}

const EXPANSION int = 1000000 - 1

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
	expansionRows, expansionCols := expansionOfSpace(string(b))
	galaxyCount := 0
	space := []SpaceNode{}

	for row, line := range strings.Split(input, "\n") {
		for col, ch := range line {
			if ch == '#' {
				galaxyCount++
				space = append(space, SpaceNode{row: row, col: col, galaxy: galaxyCount})
			}
		}
	}

	i := 0
	result := 0
	totalExpansion := 0
	for i < len(space)-1 {
		for _, node := range space[i+1:] {
			totalExpansion += expansionNodesBetween(space[i], node, expansionRows, "row")
			totalExpansion += expansionNodesBetween(space[i], node, expansionCols, "column")

			colDiff := 0
			if node.col < space[i].col {
				colDiff = space[i].col - node.col
			} else {
				colDiff = node.col - space[i].col
			}

			rowDiff := 0
			if node.row < space[i].row {
				rowDiff = space[i].row - node.row
			} else {
				rowDiff = node.row - space[i].row
			}
			distance := colDiff + rowDiff
			result += distance
		}
		i++
	}
	println(result + totalExpansion*EXPANSION)
}

func expansionNodesBetween(node_1 SpaceNode, node_2 SpaceNode, expansion []int, t string) int {
	amount := 0
	if t == "row" {

		for _, value := range expansion {
			if node_1.row > node_2.row {
				if value < node_1.row && value > node_2.row {
					amount++
				}
			} else {
				if value > node_1.row && value < node_2.row {
					amount++
				}
			}

		}
		return amount
	}
	if t == "column" {
		for _, value := range expansion {
			if node_1.col > node_2.col {
				if value < node_1.col && value > node_2.col {
					amount++
				}
			} else {
				if value > node_1.col && value < node_2.col {
					amount++
				}
			}
		}
		return amount
	}
	return 0
}

func expansionOfSpace(input string) ([]int, []int) {
	expansionRows, expansionCols := []int{}, []int{}
	rows := strings.Split(input, "\n")
	for row, line := range rows {
		if !strings.Contains(line, "#") {
			expansionRows = append(expansionRows, row)
		}

		i := 0
		expansionCol := true
		for i < len(line) && row < len(line) {
			if rows[i][row] == '#' {
				expansionCol = false
				break
			}
			i++
		}
		if expansionCol {
			expansionCols = append(expansionCols, row)
		}
	}
	return expansionRows, expansionCols

}
