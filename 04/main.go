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

var dateAndAction = regexp.MustCompile(`^\[.+:(.+)] (.+)$`)
var guardAction = regexp.MustCompile(`Guard #(.+) begins shift`)

func (g guards) addSleep(guard int, start, stop int) {
	_, ok := g[guard]
	if !ok {
		g[guard] = make(map[int]int)
	}
	// Update the counter of the minutes during which the guard was asleep
	for i := start; stop > i; i++ {
		g[guard][i]++
	}
}

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

	guards := newGuards(lines)

	guard, sleep := guards.mostAsleep()
	minute, times := guards[guard].mostFrequent()
	fmt.Println("Strategy 1:")
	fmt.Printf("Guard #%d slept the most with %d minutes of sleep.\n", guard, sleep)
	fmt.Printf("He was the most asleep during minute #%d (%d times).\n", minute, times)
	fmt.Printf("The solution to the first phase is %d * %d = %d.\n\n", guard, minute, guard*minute)

	guard, minute, count := guards.mostAsleepDuringSameMinute()
	fmt.Println("Strategy 2:")
	fmt.Printf("Minute #%d is the one during which a guard was asleep the most (%d times).\n", minute, count)
	fmt.Printf("That guard is #%d.\n", guard)
	fmt.Printf("The solution to the second phase is %d * %d = %d.\n", guard, minute, guard*minute)
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

func newGuards(lines []string) guards {
	g := make(guards)
	var currentGuard, start int

	for _, line := range lines {

		// Get the date and the action
		matches := dateAndAction.FindStringSubmatch(line)
		minute, action := matches[1], matches[2]

		// This will match if it's a new shift
		guardMatches := guardAction.FindStringSubmatch(action)

		switch {
		case len(guardMatches) > 0: // new shift
			currentGuard, _ = strconv.Atoi(guardMatches[1])
		case action == "falls asleep":
			start, _ = strconv.Atoi(minute)
		case action == "wakes up":
			stop, _ := strconv.Atoi(minute)
			g.addSleep(currentGuard, start, stop)
		}
	}

	return g
}
