package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrategy1(t *testing.T) {
	lines, err := readLines("input.txt")
	require.NoError(t, err)

	guards := newGuards(lines)

	guard, count := guards.mostAsleep()
	assert.Equal(t, 3457, guard)
	assert.Equal(t, 504, count)

	minute, count := guards[guard].mostFrequent()
	assert.Equal(t, 40, minute)
	assert.Equal(t, 14, count)
}

func TestStrategy2(t *testing.T) {
	lines, err := readLines("input.txt")
	require.NoError(t, err)

	guards := newGuards(lines)

	guard, minute, count := guards.mostAsleepDuringSameMinute()
	assert.Equal(t, 1901, guard)
	assert.Equal(t, 47, minute)
	assert.Equal(t, 19, count)
}
