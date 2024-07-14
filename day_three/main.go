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
	row    int
	col    int
	value  int
	length int
}

type Coordinate struct {
	row int
	col int
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
	var symbolCoordinates []Coordinate
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := Numbers.FindAllStringIndex(line, -1)
		symbols := Symbols.FindAllIndex([]byte(line), -1)

		for _, value := range numbers {
			machinePartValue, err := strconv.Atoi(line[value[0]:value[1]])
			if err != nil {
				log.Fatal("Something went terriby wrong :(")
			}

			machineParts = append(machineParts, MachinePart{
				row:    row,
				value:  machinePartValue,
				col:    value[0],
				length: len(line[value[0]:value[1]]),
			})
		}

		for _, value := range symbols {
			symbolCoordinates = append(symbolCoordinates, Coordinate{col: value[0], row: row})
		}
		row++
	}

	for _, part := range machineParts {
		result += machinePartCheck(part, symbolCoordinates)
	}
	println("Value: ", result)
}

func machinePartCheck(part MachinePart, symbolCoordinates []Coordinate) int {
	for i := 0; i < part.length; i++ {
		for iCol := -1; iCol <= 1; iCol++ {
			for iRow := -1; iRow <= 1; iRow++ {
				coordinate := Coordinate{col: part.col + i + iCol, row: part.row + iRow}
				if arrIncludes(coordinate, symbolCoordinates) {
					return part.value
				}
			}
		}
	}
	return 0
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

func arrIncludes(val Coordinate, from []Coordinate) bool {
	for _, valFrom := range from {
		if val.col == valFrom.col && val.row == valFrom.row {
			return true
		}
	}
	return false
}
