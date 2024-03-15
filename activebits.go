package piscine

func ActiveBits(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			count++
		}
	}
	return count - 1
}
