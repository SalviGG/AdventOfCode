package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result := dayThreeProblemOne()
	fmt.Println("The result for day three problem one is: ", result)
	result2 := dayThreeProblemTwo()
	fmt.Println("The result for day three problem two is: ", result2)
}

func dayThreeProblemOne() int {
	file, err := os.Open("dayThreeInput1.txt")
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var b bytes.Buffer
	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(b.String(), -1)
	rd := regexp.MustCompile(`\d+`)
	total := 0
	for _, value := range matches {
		numbers := rd.FindAllString(value, -1)
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		total += n1 * n2
	}
	return total
}

func dayThreeProblemTwo() int {
	file, err := os.Open("dayThreeInput1.txt")
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var b bytes.Buffer
	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}
	r := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	matches := r.FindAllString(b.String(), -1)
	rd := regexp.MustCompile(`\d+`)
	total := 0
	sumFlag := true
	for _, value := range matches {
		if value == "do()" {
			sumFlag = true
		} else if value == "don't()" {
			sumFlag = false
		}
		if sumFlag && (value != "don't()" && value != "do()") {
			numbers := rd.FindAllString(value, -1)
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])
			total += n1 * n2
		}
	}
	return total
}
