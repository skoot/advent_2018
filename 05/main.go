package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := processAll(input, -1)
	fmt.Println("Phase 1:", len(s))
	fmt.Println("Phase 2:", search(s))
}

func annihilate(r1, r2 rune) bool {
	return r1 != r2 && unicode.ToLower(r1) == unicode.ToLower(r2)
}

func processOnce(s string, ignore rune) (newString string, changes int) {
	var out []byte
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if annihilate(rune(s[i]), rune(s[i+1])) {
				i++
				changes++
				continue
			}
		}
		if unicode.ToLower(rune(s[i])) == ignore {
			changes++
			continue
		}
		out = append(out, s[i])
	}
	return string(out), changes
}

func processAll(s string, ignore rune) string {
	for {
		var changes int
		s, changes = processOnce(s, ignore)
		if changes == 0 {
			return s
		}
	}
}

func search(s string) int {
	var minLen int
	for i := 0; i < 26; i++ {
		r := rune('a' + i)
		l := len(processAll(s, r))
		if minLen == 0 || l < minLen {
			minLen = l
		}
	}

	return minLen
}
