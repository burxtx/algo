package sorting

import "testing"

func TestRecurseBinarySearch(t *testing.T) {
	tests := []struct {
		input  []int32
		key    int32
		wanted int
	}{
		{[]int32{8, 13, 20, 40}, 21, 1},
	}
	for _, test := range tests {
		res := RecurseBinarySearch(test.input, test.key)
		if res != test.wanted {
			t.Errorf("%v != %d", res, test.wanted)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input  []int32
		key    int32
		wanted int
	}{
		{[]int32{8, 13, 20, 40}, 21, 1},
	}
	for _, test := range tests {
		res := BinarySearch(test.input, test.key)
		if res != test.wanted {
			t.Errorf("%v != %d", res, test.wanted)
		}
	}
}
