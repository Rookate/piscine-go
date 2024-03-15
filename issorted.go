package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	ascending := f(a[0], a[1]) < 0

	for k := 0; k < len(a)-1; k++ {
		if (ascending && f(a[k], a[k+1]) > 0) || (!ascending && f(a[k], a[k+1]) < 0) {
			return false
		}
	}

	return true
}
