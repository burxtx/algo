package sorting

func BinarySearch(xs []int32, x int32) int {
	l := 0
	u := len(xs)
	for {
		if l < u {
			m := (l + u) / 2
			if x == xs[m] {
				return m
			} else if x < xs[m] {
				u = m
			} else if x > xs[m] {
				l = m + 1
			}
		} else {
			break
		}
	}
	return l
}

func RecurseBinarySearch(xs []int32, x int32) int {
	l := 0
	u := len(xs)
	if l < u {
		m := u / 2
		if x < xs[m] {
			u = m
			l = RecurseBinarySearch(xs[l:u], x)
		} else if x > xs[m] {
			l = m + 1
			l = RecurseBinarySearch(xs[l:u], x)
		} else if x == xs[m] {
			return m
		}
	}
	return l
}
