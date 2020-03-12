package main

import (
	"testing"
)
func TestMiniMaxSum(t *testing.T) {
	var input []int64= []int64{140638725, 436257910, 953274816, 734065819, 362748590}
	var expected []int64 = []int64{1673711044, 2486347135}
	min, max := miniMaxSum(input)
	if min != expected[0]{
		t.Error("Expected ",expected[0])
	}

	if max != expected[1]{
		t.Error("Expected ",expected[1])
	}
}