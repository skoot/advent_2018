package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/skoot/advent_2018/file"
)

func main() {
	input, err := file.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phase 1:", total(input))
	fmt.Println("Phase 2:", findDuplicate(input))
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
