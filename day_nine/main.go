package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input file>")
		os.Exit(2)
	}

	b, err := os.ReadFile(os.Args[1])

	if err != nil {
		println("Can't open file %s", os.Args[1])
		os.Exit(1)
	}

	histories := strings.Split(string(b), "\n")

	result := 0

	for _, history := range histories {
		dataSet := []int{}
		for _, dataPoint := range strings.Split(history, " ") {
			val, err := strconv.Atoi(dataPoint)
			if err != nil {
				println("Invalid input: %s", dataPoint)
				println(err)
				os.Exit(1)
			}
			dataSet = append(dataSet, val)
		}
		result += extrapolateNextValue(dataSet)
	}
	println("Result: ", result)
}

func extrapolateNextValue(dataSet []int) int {
	reducedSequences := [][]int{dataSet}

	for !readyToExtrapolate(reducedSequences[len(reducedSequences)-1]) {
		reducedSequences = append(reducedSequences, reduceSequence(reducedSequences[len(reducedSequences)-1]))
	}
	reducedSequences[len(reducedSequences)-1] = append(reducedSequences[len(reducedSequences)-1], 0)

	slices.Reverse(reducedSequences)

	current := 0
	for idx := range reducedSequences {
		if idx == len(reducedSequences)-1 {
			break
		}
		next := reducedSequences[idx+1]
		current = next[0] - current
	}

	return current
}

func reduceSequence(sequence []int) []int {
	reducedSequence := []int{}

	for i := 0; i < len(sequence)-1; i++ {
		reducedSequence = append(reducedSequence, sequence[i+1]-sequence[i])
	}
	return reducedSequence
}

func readyToExtrapolate(sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			return false
		}
	}
	return true
}
