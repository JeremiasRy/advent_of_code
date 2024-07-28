package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time   int
	record int
}

var numbers = regexp.MustCompile(`\d+.*`)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go input.txt")
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var races []Race

	scanner.Scan()
	raceTimesStr := numbers.FindAllString(scanner.Text(), -1)
	scanner.Scan()
	raceRecordsstr := numbers.FindAllString(scanner.Text(), -1)

	for idx := range raceTimesStr {
		raceTime, err := strconv.Atoi(strings.ReplaceAll(raceTimesStr[idx], " ", ""))

		if err != nil {
			log.Fatal(err)
		}

		raceRecord, err := strconv.Atoi(strings.ReplaceAll(raceRecordsstr[idx], " ", ""))

		if err != nil {
			log.Fatal(err)
		}

		races = append(races, Race{time: raceTime, record: raceRecord})
	}

	var possibilitiesArr []int

	for idx := range races {
		possibilites := 0
		for i := 0; i < races[idx].time; i++ {
			result := i * (races[idx].time - i)
			if result > races[idx].record {
				possibilites++
			}
		}
		possibilitiesArr = append(possibilitiesArr, possibilites)
	}

	for len(possibilitiesArr) != 1 {
		possibilitiesArr[0] = possibilitiesArr[0] * possibilitiesArr[1]
		possibilitiesArr = append(possibilitiesArr[:1], possibilitiesArr[2:]...)
	}

	println(possibilitiesArr[0])
}
