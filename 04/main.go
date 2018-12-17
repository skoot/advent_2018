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
	freq := make(map[interface{}]int) // map of guardID to minutes asleep

	for _, s := range ss {
		freq[s.guardID] += len(s.asleep)
	}

	k, _ := max(freq)
	return k.(int)
}

func mostMinuteAsleep(guard int, ss []shift) int {
	freq := make(map[interface{}]int) // map of minute number to number of times the guard was asleep during it

	for _, s := range ss {
		if s.guardID != guard {
			continue
		}
		for _, minute := range s.asleep {
			freq[minute]++
		}
	}

	k, _ := max(freq)
	return k.(int)
}

func mostSameMinuteAsleep(ss []shift) (int, int) {
	freq := make(map[int]map[interface{}]int) // (map of guard ID to (map of minute to (number of times the guard was asleep during it)))

	for _, s := range ss {
		c, ok := freq[s.guardID]
		if !ok {
			c = make(map[interface{}]int)
		}
		for _, minute := range s.asleep {
			c[minute]++
		}
		freq[s.guardID] = c
	}

	type guardMinute struct {
		guardID int
		minute  int
	}

	f := make(map[interface{}]int)
	for guardID, minutes := range freq {
		bestMinute, maxCount := max(minutes)
		if bestMinute == nil {
			continue
		}
		f[guardMinute{guardID, bestMinute.(int)}] = maxCount
	}

	bestGuard, _ := max(f)

	return bestGuard.(guardMinute).guardID, bestGuard.(guardMinute).minute
}
