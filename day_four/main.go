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
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		cardSplit := strings.Split(line, ":")
		onlyNumbers := cardSplit[1]
		split := strings.Split(onlyNumbers, "|")

		winningNumbersStrArr := numbers.FindAllString(split[0], -1)
		myNumbersStrArr := numbers.FindAllString(split[1], -1)

		var winningNumbers []int
		var myNumbers []int

		for _, str := range winningNumbersStrArr {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal("Failed")
			}
			winningNumbers = append(winningNumbers, num)
		}

		for _, str := range myNumbersStrArr {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal("Failed")
			}
			myNumbers = append(myNumbers, num)
		}

		cardResult := 0
		for _, num := range winningNumbers {
			for _, myNum := range myNumbers {
				if num == myNum {
					if cardResult == 0 {
						cardResult = 1
						continue
					}
					cardResult *= 2
				}
			}
		}

		result += cardResult
	}
	println("Result: ", result)
}
