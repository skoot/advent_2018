package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phase 1:", total(input))
	fmt.Println("Phase 2:", findDuplicate(input))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func total(changes []string) int {
	f := 0
	for _, change := range changes {
		f += read(change)
	}

	return f
}

func findDuplicate(changes []string) int {
	f := 0
	found := map[int]bool{f: true}

	i := 0
	for {
		f += read(changes[i])
		if found[f] {
			return f
		}
		found[f] = true
		if i == len(changes)-1 {
			i = 0
		} else {
			i++
		}
	}
}

func read(s string) int {
	if strings.HasPrefix(s, "+") {
		s = s[1:]
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return value
}
