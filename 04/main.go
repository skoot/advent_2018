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

type guards map[int]minutes

func (g guards) mostAsleep() (guard int, count int) {
	var maxMinutes int
	var maxGuard int
	for guard, minutes := range g {
		total := minutes.total()
		if total > maxMinutes {
			maxMinutes = total
			maxGuard = guard
		}
	}

	return maxGuard, maxMinutes
}

func (g guards) mostAsleepDuringSameMinute() (guard int, minute int, count int) {
	var maxGuard int
	var maxMinute int
	var maxCount int
	for guard, minutes := range g {
		minute, count := minutes.mostFrequent()
		if count > maxCount {
			maxCount = count
			maxMinute = minute
			maxGuard = guard
		}
	}

	return maxGuard, maxMinute, maxCount
}

type minutes map[int]int

func (m minutes) mostFrequent() (minute int, count int) {
	var maxMinute int
	var maxCount int
	for minute, count := range m {
		if count > maxCount {
			maxCount = count
			maxMinute = minute
		}
	}

	return maxMinute, maxCount
}

func (m minutes) total() int {
	var total int
	for _, count := range m {
		total += count
	}

	return total
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	guards := toGuards(lines)

	guard, _ := guards.mostAsleep()
	minute, _ := guards[guard].mostFrequent()
	fmt.Println("Strategy 1:", guard, minute, guard*minute)

	guard, minute, count := guards.mostAsleepDuringSameMinute()
	fmt.Println("Strategy 2:", guard, minute, count, guard*minute)
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

var dateAndAction = regexp.MustCompile(`^\[.+:(.+)] (.+)$`)
var guardAction = regexp.MustCompile(`Guard #(.+) begins shift`)

func toGuards(lines []string) guards {
	asleep := make(map[int]minutes)
	var guard int
	var start int

	for _, line := range lines {

		// Get the date and the action
		matches := dateAndAction.FindStringSubmatch(line)
		minute, action := matches[1], matches[2]

		// This will match if it's a new shift
		guardMatches := guardAction.FindStringSubmatch(action)

		switch {
		case len(guardMatches) > 0: // new shift
			guard, _ = strconv.Atoi(guardMatches[1])
		case action == "falls asleep":
			start, _ = strconv.Atoi(minute)
		case action == "wakes up":
			stop, _ := strconv.Atoi(minute)
			_, ok := asleep[guard]
			if !ok {
				asleep[guard] = make(map[int]int)
			}
			// Update the counter of the minutes during which the guard was asleep
			for i := start; stop > i; i++ {
				asleep[guard][i]++
			}
		}
	}

	return asleep
}
