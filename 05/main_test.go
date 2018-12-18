package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_processOnce(t *testing.T) {
	in := "abBcD"
	out, changes := processOnce(in)
	assert.Equal(t, "acD", out)
	assert.Equal(t, 1, changes)
}

func Test_processOnce_complex(t *testing.T) {
	in := "abBcDdC"
	out, changes := processOnce(in)
	assert.Equal(t, "acC", out)
	assert.Equal(t, 2, changes)
}

func Test_processAll(t *testing.T) {
	in := "dabAcCaCBAcCcaDA"
	out := processAll(in)
	assert.Equal(t, "dabCBAcaDA", out)
}
