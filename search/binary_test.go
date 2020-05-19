package search

import (
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input  float64
		wanted float64
	}{
		{2, math.Sqrt(2)},
	}
	for _, test := range tests {
		actual, err := BinarySearch(test.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != test.wanted {
			t.Errorf("%v != %f", actual, test.wanted)
		}
	}
}
