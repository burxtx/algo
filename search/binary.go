package search

import (
	"errors"
	"fmt"
	"math"
)

func BinarySearch(x float64) (float64, error) {
	var mid float64
	if x < 0 {
		return 0, errors.New("negative is not allowed")
	} else if x == 1 {
		return 1, nil
	} else if x > 1 {
		low := float64(1)
		high := x
		mid = low + ((high - low) / 2)
		for math.Abs((mid*mid - x)) > 0.00001 {
			fmt.Println(mid, low, high)
			if mid*mid < x {
				fmt.Println("low is", mid, low, high)
				low = mid
			} else if mid*mid > x {
				fmt.Println("high is", mid, low, high)
				high = mid
			}
			mid = low + ((high - low) / 2)
		}
	}
	return mid, nil
}
