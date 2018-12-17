package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type shift struct {
	guardID int
	asleep  []int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ss, err := shifts(lines)
	if err != nil {
		log.Fatal(err)
	}

	guard := mostAsleepGuard(ss)
	minute := mostMinuteAsleep(guard, ss)
	fmt.Println("Strategy 1: ", guard*minute)

	guard, minute = mostSameMinuteAsleep(ss)
	fmt.Println("Strategy 2: ", guard*minute)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	return lines, scanner.Err()
}

var dateAndAction = regexp.MustCompile(`^\[.+ .+:(.+)] (.+)$`)
var guardAction = regexp.MustCompile(`Guard #(.+) begins shift`)

func shifts(lines []string) ([]shift, error) {
	out := make([]shift, 0)
	var s shift
	var start, stop int

	for _, line := range lines {

		// Get the date and the action
		matches := dateAndAction.FindStringSubmatch(line)
		minute, action := matches[1], matches[2]

		// This will match if it's a new shift
		guardMatches := guardAction.FindStringSubmatch(action)

		switch {
		case len(guardMatches) > 0: // new shift
			if s.guardID != 0 {
				out = append(out, s)
			}
			guardID, _ := strconv.Atoi(guardMatches[1])
			s = shift{
				guardID: guardID,
			}
		case action == "falls asleep":
			start, _ = strconv.Atoi(minute)
		case action == "wakes up":
			stop, _ = strconv.Atoi(minute)
			// Fill the asleep slice with the minutes during which the guard was asleep
			for i := 0; stop > start; i++ {
				s.asleep = append(s.asleep, start)
				start++
			}
		}
	}

	return out, nil
}

func mostAsleepGuard(ss []shift) int {
	freq := make(map[int]int) // map of guardID to minutes asleep

	for _, s := range ss {
		freq[s.guardID] += len(s.asleep)
	}

	return max(freq)
}

func mostMinuteAsleep(guard int, ss []shift) int {
	freq := make(map[int]int) // map of minute number to number of times the guard was asleep during it

	for _, s := range ss {
		if s.guardID != guard {
			continue
		}
		for _, minute := range s.asleep {
			freq[minute]++
		}
	}

	return max(freq)
}

// max returns the key of the provided map that has the highest value
func max(freq map[int]int) int {
	var max, out int
	for item, count := range freq {
		if count > max {
			max = count
			out = item
		}
	}
	return out
}

type stat struct {
	guardID int
	minute  int
	count   int
}

func mostSameMinuteAsleep(ss []shift) (int, int) {
	freq := make(map[int]map[int]int) // (map of guard ID to (map of minute to (number of times the guard was asleep during it)))

	for _, s := range ss {
		c, ok := freq[s.guardID]
		if !ok {
			c = make(map[int]int)
		}
		for _, minute := range s.asleep {
			c[minute]++
		}
		freq[s.guardID] = c
	}

	stats := make([]stat, 0, len(freq))
	for guardID, f := range freq {
		var max int
		var out int
		for minute, count := range f {
			if count > max {
				max = count
				out = minute
			}
		}
		stats = append(stats, stat{
			guardID: guardID,
			minute:  out,
			count:   max,
		})
	}

	var max int
	var out stat
	for _, s := range stats {
		if s.count > max {
			max = s.count
			out = s
		}
	}

	return out.guardID, out.minute
}
