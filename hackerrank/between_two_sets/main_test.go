package main

import (
	"testing"
)

func TestSomething(t *testing.T) {
	parameters := []struct {
		a, b     []int32
		expected int32
	}{
		{[]int32{2, 4}, []int32{16, 32, 96}, 3},
		{[]int32{2, 6}, []int32{24, 36}, 2},
	}

	for i := range parameters {
		result := getTotalX(parameters[i].a, parameters[i].b)
		if result != parameters[i].expected {
			t.Error("Expected: ", parameters[i].expected, "and was ", result)

		}
	}
}
