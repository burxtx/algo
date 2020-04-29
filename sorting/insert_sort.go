package sorting

func InsertSort(xs []int32) []int32 {
	for i := 1; i < len(xs); i++ {
		x := xs[i]
		j := i - 1
		for ; j >= 0; j-- {
			if j >= 0 && x < xs[j] {
				xs[j+1] = xs[j]
			} else {
				break
			}
		}
		xs[j+1] = x
	}
	return xs
}
