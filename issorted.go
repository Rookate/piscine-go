package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	ascending := f(a[0], a[1]) < 0

	for i := 0; i < len(a)-1; i++ {
		if (ascending && f(a[i], a[i+1]) > 0) || (!ascending && f(a[i], a[i+1]) < 0) {
			return false
		}
	}

	return true
}
