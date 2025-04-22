package utility

import (
	"bufio"
	"log"
	"os"
)

func FileToStrSlice(path string) []string {
	var out []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("fail: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}
	return out
}
