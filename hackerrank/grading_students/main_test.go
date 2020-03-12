package main

import (
	"log"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	parameters := []struct {
		input, expected []int32
	}{
		{[]int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
	}

	for i := range parameters {
		log.Println("Round", i)
		result := gradingStudents(parameters[i].input)
		expected := parameters[i].expected

		if len(result) != len(expected) {
			t.Error("Size of result and expected arrays is different.")
		}
		for x := range result {
			if result[x] != expected[x] {
				t.Error("Expected: ", expected[x], "and give", result[x])
			}
		}
	}
}
