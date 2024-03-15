package piscine

func Unmatch(a []int) int {
	for _, n := range a {
		ch := 0
		for _, nb := range a {
			if nb == n {
				ch++
			}
		}
		if ch == 1 || ch%2 == 1 {
			return n
		}
	}
	return -1
}
