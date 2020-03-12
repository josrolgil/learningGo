package main

import (
	"testing"
)

func TestSomething(t *testing.T) {
	parameters := []struct {
		x1, v1, x2, v2 int32
		expected       string
	}{
		//{0, 3, 4, 2, "YES"},
		//{0, 2, 5, 3, "NO"},
		//{2, 5, 100, 1, "NO"},
		{63, 8, 94, 3, "NO"},
		{43, 2, 70, 2, "NO"},
	}

	for i := range parameters {
		result := kangaroo(parameters[i].x1, parameters[i].v1, parameters[i].x2, parameters[i].v2)
		if result != parameters[i].expected {
			t.Error("Expected: ", parameters[i].expected, "and was ", result)

		}
	}
}
