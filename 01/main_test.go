package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_total(t *testing.T) {
	input := []string{"+1", "-2", "+3", "+1"}
	f, err := total(input)
	require.NoError(t, err)
	assert.Equal(t, 3, f)
}

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name    string
		changes []string
		want    int
	}{
		{
			name:    "simple",
			changes: []string{"+1", "-1"},
			want:    0,
		},
		{
			name:    "10",
			changes: []string{"+3", "+3", "+4", "-2", "-4"},
			want:    10,
		},
		{
			name:    "5",
			changes: []string{"-6", "+3", "+8", "+5", "-6"},
			want:    5,
		},
		{
			name:    "14",
			changes: []string{"+7", "+7", "-2", "-7", "-4"},
			want:    14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findDuplicate(tt.changes)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
