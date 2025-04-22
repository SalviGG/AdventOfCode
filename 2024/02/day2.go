package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type GlobalReport struct {
	Reports   []SafetyReport
	TotalSafe int
}

type SafetyReport struct {
	Safety bool
	Rows   []IndividualState
}

type IndividualState struct {
	Value int
	State bool
}

func dayTwoProblemOne() {
	var sSlice []string
	file, err := os.Open("dayTwoInput1.txt")
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sSlice = append(sSlice, line)
	}
	for _, value := range sSlice {
		split := strings.Split(value, " ")
		var row []int
	}
}

func validRow(row []int) bool {
	v := false
	for i := 0; i < len(row); i++ {

	}
	return v

}
