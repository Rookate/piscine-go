package piscine

func Map(f func(int) bool, a []int) (res []bool) {
	for i := range a {
		res = append(res, f(a[i]))
	}
	return res
}
