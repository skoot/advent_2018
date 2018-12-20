package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distance(t *testing.T) {
	tests := []struct {
		name   string
		coord1 coord
		coord2 coord
		want   int
	}{
		{
			name:   "identical",
			coord1: coord{3, 3},
			coord2: coord{3, 3},
			want:   0,
		},
		{
			name:   "apart",
			coord1: coord{1, 3},
			coord2: coord{3, 4},
			want:   3,
		},
		{
			name:   "reversed",
			coord1: coord{3, 4},
			coord2: coord{1, 3},
			want:   3,
		},
		{
			name:   "diagonal",
			coord1: coord{4, 3},
			coord2: coord{1, 2},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, distance(tt.coord1, tt.coord2))
		})
	}
}

func Test_closest_equidistant(t *testing.T) {
	c := closestPoint(coord{1, 1}, []coord{{1, 0}, {1, 2}})
	assert.Equal(t, coord{0, 0}, c)
}

func Test_closest_closest(t *testing.T) {
	c := closestPoint(coord{1, 1}, []coord{{1, 0}, {1, 3}})
	assert.Equal(t, coord{1, 0}, c)
}

func Test_closest_diag(t *testing.T) {
	c := closestPoint(coord{2, 2}, []coord{{1, 1}, {3, 4}})
	assert.Equal(t, coord{1, 1}, c)
}

func Test_largest(t *testing.T) {
	c := []coord{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}

	fmt.Println(c)

	sizes := getSizes(c)
	fmt.Println(sizes)

	assert.Equal(t, 17, largestArea(sizes, c))
}

func Test_limits(t *testing.T) {
	c := []coord{
		{1, 1},
		{1, 6},
		{8, 3},
		{3, 4},
		{5, 5},
		{8, 9},
	}

	minX, maxX, minY, maxY := limits(c)
	assert.Equal(t, 1, minX)
	assert.Equal(t, 8, maxX)
	assert.Equal(t, 1, minY)
	assert.Equal(t, 9, maxY)
}
