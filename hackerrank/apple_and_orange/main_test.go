package main

import (
	"log"
	"testing"
)

//func calculateNumberInSpace(start int32, end int32, items []int32) int32
func TestCalculateNumberInSpace(t *testing.T) {
	parameters := []struct {
		start, end int32
		items      []int32
		expected   int32
	}{
		{4, 6, []int32{6, 7, 0}, 1},
	}
	for i := range parameters {
		result := calculateNumberInSpace(parameters[i].start, parameters[i].end, parameters[i].items)
		if result != parameters[i].expected {
			t.Error("Expected: ", parameters[i].expected, "and was", result)
		}
	}
}
func TestProcessPosition(t *testing.T) {
	parameters := []struct {
		reference           int32
		movements, expected []int32
	}{
		{4, []int32{2, 3, -4}, []int32{6, 7, 0}},
	}

	for i := range parameters {
		result := processPosition(parameters[i].reference, parameters[i].movements)
		for x := range result {
			if result[x] != parameters[i].expected[x] {
				t.Error("Expected: ", parameters[i].expected[x], "and was", result[x])
			}
		}
	}
}

func TestProcessApplesAndOranges(t *testing.T) {
	parameters := []struct {
		s, t, a, b                      int32
		apples, oranges                 []int32
		expectedApples, expectedOranges int32
	}{
		{7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4}, 1, 2},
		{7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6}, 1, 1},
	}

	for i := range parameters {
		log.Println("Round", i)
		resultApple, resultOrange := processApplesAndOranges(parameters[i].s, parameters[i].t, parameters[i].a, parameters[i].b, parameters[i].apples, parameters[i].oranges)

		if parameters[i].expectedApples != resultApple || parameters[i].expectedOranges != resultOrange {
			t.Error("Wrong result.")
		}
	}
}
