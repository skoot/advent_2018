package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/skoot/advent_2018/file"
)

func main() {
	lines, err := file.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	max2 := countDups(lines, 2)
	max3 := countDups(lines, 3)

	fmt.Println("Phase 1:", max2, max3, max2*max3)
	fmt.Println("Phase 2:", searchOneLetterDiff(lines))
}

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func order(s string) string {
	b := sortBytes(s)
	sort.Sort(b)
	return string(b)
}

func findDuplicates(s string, nb int) int {
	var nbDups int
	var current rune
	count := 1
	for _, r := range s {
		if r == current {
			count++
		} else {
			if count == nb {
				nbDups++
			}
			count = 1
		}

		current = r
	}

	// handle the case where a correct number of duplicate would be at the end of the string
	if count == nb {
		nbDups++
	}

	return nbDups
}

func countDups(lines []string, nb int) int {
	var total int
	for _, line := range lines {
		n := findDuplicates(order(line), nb)
		if n > 0 {
			total++
		}
	}

	return total
}

func oneLetterDiff(l1, l2 string) string {
	var count int
	var out string
	for i := range l1 {
		if l1[i] != l2[i] {
			count++
			out = l1[0:i] + l1[i+1:]
		}
	}

	if count != 1 {
		return ""
	}

	return out
}

func searchOneLetterDiff(lines []string) string {
	for _, line := range lines {
		for _, l := range lines {
			out := oneLetterDiff(line, l)
			if out != "" {
				return out
			}
		}
	}

	return ""
}
