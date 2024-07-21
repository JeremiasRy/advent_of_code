package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	numbers = regexp.MustCompile(`\d+`)
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go input.txt")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("Can't open file %v", err)
	}

	scanner := bufio.NewScanner(file)
	row := 0
	result := 0
	var copies []int
	for scanner.Scan() {
		if row >= len(copies) {
			copies = append(copies, 0)
		}
		copies[row] += 1 // always add the original copy

		line := scanner.Text()
		cardSplit := strings.Split(line, ":")[1]
		split := strings.Split(cardSplit, "|")

		winningNumbersStrArr := numbers.FindAllString(split[0], -1)
		myNumbersStrArr := numbers.FindAllString(split[1], -1)

		winningNumbers := parseNumbers(winningNumbersStrArr)
		myNumbers := parseNumbers(myNumbersStrArr)

		cardResult := 0
		for _, num := range winningNumbers {
			for _, myNum := range myNumbers {
				if num == myNum {
					cardResult += 1
				}
			}
		}
		if cardResult == 0 {
			row++
			continue
		}
		if len(copies)-row <= cardResult {
			for i := row + 1; i < row+1+cardResult; i++ {
				if i >= len(copies) {
					copies = append(copies, copies[row])
					continue
				}
				copies[i] += copies[row]
			}
		} else {
			for i := 0; i < cardResult; i++ {
				copies[i+row+1] += copies[row]
			}
		}
		row++
	}
	for _, scratchCards := range copies {
		result += scratchCards
	}

	println("Result: ", result)
}

func parseNumbers(strArr []string) []int {
	var result []int
	for idx := range strArr {
		num, err := strconv.Atoi(strArr[idx])
		if err != nil {
			log.Fatal("Failed")
		}
		result = append(result, num)
	}
	return result
}
