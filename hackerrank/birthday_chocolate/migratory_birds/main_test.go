package main

import (
	"testing"
)

func TestMigration(t *testing.T) {
	//func migratoryBirds(arr []int32) int32 {
	parameters := []struct {
		arr      []int32
		expected int32
	}{
		{[]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}, 3},
	}

	for i := 0; i < len(parameters); i++ {
		result := migratoryBirds(parameters[i].arr)
		if result != parameters[i].expected {
			t.Error("ERROR")
		}
	}
}
