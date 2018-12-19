package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorder(t *testing.T) {
	assert.Equal(t, "abcde", order("cbaed"))
}

func Test_findDuplicates(t *testing.T) {
	assert.Equal(t, 2, findDuplicates("aaabbcdee", 2))
}

func Test_oneLetterDiff_found(t *testing.T) {
	assert.Equal(t, "abdef", oneLetterDiff("abcdef", "abgdef"))
}

func Test_oneLetterDiff_no_diff(t *testing.T) {
	assert.Equal(t, "", oneLetterDiff("abcdef", "abcdef"))
}

func Test_oneLetterDiff_more_diff(t *testing.T) {
	assert.Equal(t, "", oneLetterDiff("abcdef", "abgdey"))
}

func Test_searchOneLetterDiff_found(t *testing.T) {
	lines := []string{"acbdef", "abcdef", "abcgef", "xcbdeg"}
	assert.Equal(t, "abcef", searchOneLetterDiff(lines))
}
