package search

func BinarySearch(x float64) (float64, error) {
	if x < 0 {
		return 0, error.New("negative is not allowed")
	} else {
		low := 0
		high := x
		for low < high && (mid**2-x) > 0.000001 {
			mid := low + ((high - low) >> 1)
			if mid**2 < x {
				low = mid
			} else if mid**2 > x {
				high = mid
			}
		}
	}
}
