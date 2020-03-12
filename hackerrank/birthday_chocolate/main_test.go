package main

import (
	"testing"
)

func TestBirthdayChocolate(t *testing.T) {
	parameters := []struct {
		s              []int32
		d, m, expected int32
	}{
		/*{[]int32{1, 1, 1, 1, 1, 1}, 3, 2, 0},
		{[]int32{1, 2, 1, 3, 2}, 3, 2, 2},
		{[]int32{4}, 4, 1, 1},*/
		{[]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, 18, 7, 3},
	}

	for i := 0; i < len(parameters); i++ {
		result := birthday(parameters[i].s, parameters[i].d, parameters[i].m)
		if result != parameters[i].expected {
			t.Error("ERROR")
		}

	}
}
