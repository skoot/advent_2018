package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/skoot/advent_2018/file"
)

func main() {
	lines, err := file.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := coords(lines)
	sizes := getSizes(c)

	fmt.Println(limits(c))

	fmt.Println("Phase 1:", largestArea(sizes, c))
	fmt.Println("Phase 2:", largestRegionWithTotalDistanceLessThan(c, 10000))
}

type coord struct {
	X int
	Y int
}

func coords(lines []string) []coord {
	out := make([]coord, len(lines))
	for i := range lines {
		out[i] = parseLine(lines[i])
	}

	return out
}

func parseLine(line string) coord {
	parts := strings.Split(line, ", ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coord{
		X: x,
		Y: y,
	}
}

func distance(coord1, coord2 coord) int {
	return abs(coord1.X-coord2.X) + abs(coord1.Y-coord2.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func limits(coords []coord) (minX, maxX, minY, maxY int) {
	for _, c := range coords {
		if minX == 0 || c.X < minX {
			minX = c.X
		}
		if maxX == 0 || c.X > maxX {
			maxX = c.X
		}
		if minY == 0 || c.Y < minY {
			minY = c.Y
		}
		if maxY == 0 || c.Y > maxY {
			maxY = c.Y
		}
	}

	return
}

func getSizes(coords []coord) map[coord]int {
	out := make(map[coord]int)
	minX, maxX, minY, maxY := limits(coords)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			c := coord{x, y}
			out[closestPoint(c, coords)]++
		}
	}

	return out
}

func closestPoint(c coord, coords []coord) coord {
	var min int
	var out coord
	for _, cc := range coords {
		d := distance(c, cc)
		switch {
		case c == cc: // there's a point here
			return c
		case min == d: // this one is the same distance as another one
			min = 0
			out = coord{0, 0}
		case min == 0 || d < min:
			min = d
			out = cc
		}
	}

	return out
}

func largestArea(sizes map[coord]int, coords []coord) int {
	var largest int
	minX, maxX, minY, maxY := limits(coords)
	for key, value := range sizes {
		if key.X == 0 || key.Y == 0 || key.X == minX || key.X == maxX || key.Y == minY || key.Y == maxY {
			continue
		}
		if value > largest {
			largest = value
		}
	}

	return largest
}

func totalDistances(c coord, coords []coord) int {
	var out int
	for _, cc := range coords {
		out += distance(c, cc)
	}

	return out
}

func largestRegionWithTotalDistanceLessThan(coords []coord, maxDistance int) int {
	var out int
	minX, maxX, minY, maxY := limits(coords)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			c := coord{x, y}
			td := totalDistances(c, coords)
			if td < maxDistance {
				out++
			}
		}
	}

	return out
}
