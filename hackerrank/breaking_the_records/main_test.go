package main

import (
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	parameters := []struct {
		scores                   []int32
		expectedMax, expectedMin int32
	}{
		{[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, 2, 4},
		{[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, 4, 0},
	}

	for i := range parameters {
		result := breakingRecords(parameters[i].scores)
		if result[0] != parameters[i].expectedMax {
			t.Error("Expected max: ", parameters[i].expectedMax, "and was ", result[0])
		}
		if result[1] != parameters[i].expectedMin {
			t.Error("Expected min: ", parameters[i].expectedMin, "and was ", result[1])
		}
	}
}
