package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	dialLeftPointingAtZeroCount := 0
	dialPointingAtZeroCount := 0

	dialPoint := 50

	for line := range strings.SplitSeq(string(input), "\n") {
		rotationsNumber, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting number", err)
		}

		if line[:1] == "L" {
			rotationsNumber *= -1
		}
		dialPoint += rotationsNumber

		if dialPoint%100 == 0 {
			dialLeftPointingAtZeroCount++
		}

		if dialPoint >= 100 {
			dialPointingAtZeroCount += dialPoint / 100
			dialPoint %= 100
		} else if dialPoint < 0 {
			if dialPoint-rotationsNumber != 0 {
				dialPointingAtZeroCount++
			}

			dialPointingAtZeroCount -= dialPoint / 100

			dialPoint %= 100
			if dialPoint < 0 {
				dialPoint += 100
			}

		} else if dialPoint == 0 {
			dialPointingAtZeroCount++
		}

	}
	fmt.Println("First Puzzle Answer", dialLeftPointingAtZeroCount)
	fmt.Println("Second Puzzle Answer", dialPointingAtZeroCount)
}
