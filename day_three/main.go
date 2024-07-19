package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	Numbers = regexp.MustCompile(`\d+`)
	Symbols = regexp.MustCompile(`[^\d.]`)
)

type MachinePart struct {
	positions []Coordinate
	value     int
}

type Coordinate struct {
	row int
	col int
}

type Symbol struct {
	ch                   string
	position             Coordinate
	machinePartsAdjacent []MachinePart
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file to read, please specify.")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Can't find file ", os.Args[1])
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	row := 0
	var machineParts []MachinePart
	var symbols []Symbol
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbersIdx := Numbers.FindAllStringIndex(line, -1)
		symbolsIdx := Symbols.FindAllIndex([]byte(line), -1)

		for _, value := range numbersIdx {
			machinePartValue, err := strconv.Atoi(line[value[0]:value[1]])
			if err != nil {
				log.Fatal("Something went terriby wrong :(")
			}

			var machinePartPositions []Coordinate
			for i := 0; i < len(line[value[0]:value[1]]); i++ {
				for iCol := -1; iCol <= 1; iCol++ {
					for iRow := -1; iRow <= 1; iRow++ {
						machinePartPositions = append(machinePartPositions, Coordinate{row: row + iRow, col: i + value[0] + iCol})
					}
				}
			}
			machineParts = append(machineParts, MachinePart{
				positions: machinePartPositions,
				value:     machinePartValue,
			})
		}

		for _, value := range symbolsIdx {
			symbol := Symbol{position: Coordinate{col: value[0], row: row}, ch: string(line[value[0]])}
			symbols = append(symbols, symbol)
		}
		row++
	}
	found := false
	for _, part := range machineParts {
		found = false
		for j := range symbols {
			if found {
				continue
			}
			for _, partPosition := range part.positions {
				if symbols[j].position.col == partPosition.col && symbols[j].position.row == partPosition.row && !found {
					symbols[j].machinePartsAdjacent = append(symbols[j].machinePartsAdjacent, part)
					found = true
				}
			}
		}
	}

	for _, symbol := range symbols {
		if len(symbol.machinePartsAdjacent) == 2 && symbol.ch == "*" {
			result += symbol.machinePartsAdjacent[0].value * symbol.machinePartsAdjacent[1].value
			continue
		}
	}

	println("Result: ", result)
}

func arrRange(start int, len int) []int {
	arr := make([]int, len)

	i := 0
	for i < len {
		arr[i] = start + i
		i++
	}

	return arr
}
