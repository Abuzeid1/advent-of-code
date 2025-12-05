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
	DialLeftPointingAtZeroCount := 0
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
			DialLeftPointingAtZeroCount++
		}

	}
	fmt.Println("First Puzzle Answer", DialLeftPointingAtZeroCount)

}
