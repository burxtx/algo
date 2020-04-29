package sorting

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []struct {
		input  []int32
		wanted string
	}{
		{[]int32{40, 8, 13, 20}, "8,13,20,40"},
		{[]int32{4, 5, 6, 1, 3, 2}, "1,2,3,4,5,6"},
	}
	for _, test := range tests {
		res := InsertSort(test.input)

		if res[0] != 20 {
			t.Errorf("%v != %s", res, test.wanted)
		}
	}
}
