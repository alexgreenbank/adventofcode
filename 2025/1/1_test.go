package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_doMove(t *testing.T) {
	var tests = []struct {
		testName     string
		curr         int
		move         string
		expectedDial int
		expectedP1   int
		expectedP2   int
	}{
		// Simple cases
		{"R1 from 50", 50, "R1", 51, 0, 0},
		{"L1 from 50", 50, "L1", 49, 0, 0},
		{"R49 from 50", 50, "R49", 99, 0, 0},
		{"L49 from 50", 50, "L49", 1, 0, 0},
		// Ending on 0
		{"R50 from 50", 50, "R50", 0, 1, 1},
		{"L50 from 50", 50, "L50", 0, 1, 1},
		// Passing 0
		{"R51 from 50", 50, "R51", 1, 0, 1},
		{"L51 from 50", 50, "L51", 99, 0, 1},
		// Ending on 0 with an extra full spin
		{"R150 from 50", 50, "R150", 0, 1, 2},
		{"L150 from 50", 50, "L150", 0, 1, 2},
		// Passing 0 with an extra full spin
		{"R151 from 50", 50, "R151", 1, 0, 2},
		{"L151 from 50", 50, "L151", 99, 0, 2},
		// Let's test some edge cases starting on 0
		{"R1 from 0", 0, "R1", 1, 0, 0},
		{"L1 from 0", 0, "L1", 99, 0, 0},
		{"R100 from 0", 0, "R100", 0, 1, 1},
		{"L100 from 0", 0, "L100", 0, 1, 1},
		{"R200 from 0", 0, "R200", 0, 1, 2},
		{"L200 from 0", 0, "L200", 0, 1, 2},
		// {"Fatal", -1, "C100", 99, 0, 2},
		// {"Fatal", 50, "C", 99, 0, 2},
		// {"Fatal", 50, "C100", 99, 0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			dial, p1inc, p2inc := doMove(tt.curr, tt.move)
			require.Equal(t, tt.expectedDial, dial)
			require.Equal(t, tt.expectedP1, p1inc)
			require.Equal(t, tt.expectedP2, p2inc)
		})
	}

}
