package main

import (
	"testing"
)

func TestDivisibleSumPairs(t *testing.T) {
	//func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	parameters := []struct {
		n, k     int32
		ar       []int32
		expected int32
	}{
		{6, 3, []int32{1, 3, 2, 6, 1, 2}, 5},
	}
	for i := 0; i < len(parameters); i++ {
		result := divisibleSumPairs(parameters[i].n, parameters[i].k, parameters[i].ar)
		if result != parameters[i].expected {
			t.Error("ERROR")
		}
	}
}
