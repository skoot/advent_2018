package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toGuards(t *testing.T) {
	lines := []string{
		`[x:00] Guard #1 begins shift`,
		`[x:10] falls asleep`,
		`[x:12] wakes up`,
		`[x:09] Guard #2 begins shift`,
		`[x:20] falls asleep`,
		`[x:22] wakes up`,
		`[x:53] Guard #1 begins shift`,
		`[x:10] falls asleep`,
		`[x:13] wakes up`,
	}
	g := toGuards(lines)
	assert.Equal(t, guards{
		1: []int{10, 11, 10, 11, 12},
		2: []int{20, 21},
	}, g)
}

func Test_guards_mostAsleep(t *testing.T) {
	g := guards{
		1: []int{10, 11, 10, 11, 12},
		2: []int{20, 21},
	}

	guardID, nbMinutes := g.mostAsleep()
	assert.Equal(t, 1, guardID)
	assert.Equal(t, 5, nbMinutes)
}
