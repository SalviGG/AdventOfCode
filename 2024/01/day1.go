package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func problemOne() int {
	c := readFile("dayOneInput1.txt")
	c2 := readFile("dayTwoInput2.txt")

	sort.Ints(c)
	sort.Ints(c2)

	var s []int

	for i := 0; i < len(c); i++ {
		if c[i] > c2[i] {
			s = append(s, (c[i] - c2[i]))
		} else {
			s = append(s, (c2[i] - c[i]))
		}
	}

	var sum int
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}

	return sum

}

func readFile(filename string) []int {
	var intSlice []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		intSlice = append(intSlice, n)
	}

	return intSlice
}

func problemTwo() {
	c := readFile("dayOneInput1.txt")
	c2 := readFile("dayTwoInput2.txt")

	sort.Ints(c)
	sort.Ints(c2)
	var count []int

	for i := 0; i < len(c); i++ {
		sum := 0
		for f := 0; f < len(c2); f++ {
			if c[i] == c2[f] {
				sum++
			}
		}
		count = append(count, (c[i] * sum))
	}

	var totalSum int
	fmt.Println(count)

	for i := 0; i < len(count); i++ {
		totalSum += count[i]
	}
	fmt.Printf("The result of problem two is: %v", totalSum)
}
