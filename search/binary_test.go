package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input  float64
		wanted float64
	}{
		{2, 1.414},
	}
	for _, test := range tests {
		actual, err := BinarySearch(test.input)
		if err != nil {
			t.Errorf(err.Error())
		}
		if actual != test.wanted {
			t.Errorf("%v != %s", actual, test.wanted)
		}
	}
}
