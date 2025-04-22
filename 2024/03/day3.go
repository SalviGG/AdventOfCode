package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func dayThreeProblemOne() {
	file, err := os.Open("dayTwoInput1.txt")
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var b bytes.Buffer
	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}
}
