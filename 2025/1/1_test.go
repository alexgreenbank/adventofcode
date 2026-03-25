package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_doMove(t *testing.T) {
	var tests = []struct {
		testName string
		curr int
		move string
		expectedDial int
		expectedP1 int
	}{
		{"R1 from 50", 50, "R1", 51, 0},
		{"L1 from 50", 50, "L1", 49, 0},
		{"R49 from 50", 50, "R49", 99, 0},
		{"L49 from 50", 50, "L49", 1, 0},
		{"R50 from 50", 50, "R50", 0, 1},
		{"L50 from 50", 50, "L50", 0, 1},
		{"R51 from 50", 50, "R51", 1, 0},
		{"L51 from 50", 50, "L51", 99, 0},
		{"R150 from 50", 50, "R150", 0, 1},
		{"L150 from 50", 50, "L150", 0, 1},
		{"R151 from 50", 50, "R151", 1, 0},
		{"L151 from 50", 50, "L151", 99, 0},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			dial, p1inc := doMove(tt.curr, tt.move)
			require.Equal(t, dial, tt.expectedDial)
			require.Equal(t, p1inc, tt.expectedP1)
		})
	}

}
