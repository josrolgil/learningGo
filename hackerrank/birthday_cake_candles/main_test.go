package main

import "testing"

func TestBirthdayCakeCandles(t *testing.T){
	var input []int32 = []int32{3, 2 ,1 ,3}
	var expected int32 = 2

	if birthdayCakeCandles(input) != expected{
		t.Error("Expected ",expected)
	}
}