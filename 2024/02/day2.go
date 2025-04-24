package main

import (
	"AdventOfCode2024/utility"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputRows := utility.FileToStrSlice("./dayTwoInput.txt")
	solutionProblemOne := problemOne(inputRows)
	fmt.Println("The number of safety reports for day 2 problem 1 is: " + strconv.Itoa(solutionProblemOne))
}

func problemOne(rows []string) int {
	sr := 0
	for _, value := range rows {
		split := strings.Split(value, " ")
		if validRow(split) {
			sr++
		}
	}
	return sr
}

func validRow(row []string) bool {
	v := true
	asc, desc := false, false

	for i := 1; i < len(row); i++ {
		current, _ := strconv.Atoi(row[i])
		previous, _ := strconv.Atoi(row[i-1])
		diff := current - previous
		if diff > 0 {
			desc = true
		} else if diff < 0 {
			asc = true
		} else {
			return false
		}
		if desc && asc {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return v
}

// func problemTwo(row []string) int {
// }
