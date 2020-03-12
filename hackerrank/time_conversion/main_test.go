package main

import (
	"log"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	parameters := []struct {
		input, expected string
	}{
		{"07:05:45PM", "19:05:45"}, {"07:05:45AM", "07:05:45"},
	}

	for i := range parameters {
		log.Println("Round", i)
		if timeConversion(parameters[i].input) != parameters[i].expected {
			t.Error("Expected: ", parameters[i].expected)
		}
	}
}
