package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sumDivs(t *testing.T) {
	var tests = []struct {
		testName string
		n1       int
		n2       int
		div      int
		expected int
	}{
		// Simple cases
		{"Single match at start of range", 11, 21, 11, 11},
		{"Single match at end of range", 12, 22, 11, 22},
		{"Match at start and end of range", 11, 22, 11, 33},
		{"No match in range", 12, 21, 11, 0},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			r := sumDivs(tt.n1, tt.n2, tt.div)
			require.Equal(t, tt.expected, r)
		})
	}
}

func Test_doRange(t *testing.T) {
	var tests = []struct {
		testName string
		n1       int
		n2       int
		expected int
	}{
		// Simple cases
		{"Single match at start of range", 11, 21, 11},
		{"Single match at end of range", 12, 22, 22},
		{"Match at start and end of range", 11, 22, 33},
		{"No match in range", 12, 21, 0},
		// Span two digit lengths
		{"Two digit lengths no matches", 5, 10, 0},
		{"Two digit lengths single match", 5, 21, 11},
		{"Two digit lengths multiple matches", 5, 22, 33},
		// Span three digit lengths
		{"Three digit lengths multiple matches", 5, 105, 495}, // sum(1,9) = 45, 45 * 101 = 495
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			r := doRange(tt.n1, tt.n2)
			require.Equal(t, tt.expected, r)
		})
	}
}
