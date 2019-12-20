package sorting

func InsertSort(xs []int32) []int32 {
	for i := 1; i < len(xs); i++ {
		x := xs[i]
		j := i - 1
		for {
			if j >= 0 && x < xs[j] {
				xs[j+1] = xs[j]
				j = j - 1
			} else {
				xs[j+1] = x
				break
			}
		}
	}
	return xs
}

func InsertSort() {

}

func Insert() {

}
