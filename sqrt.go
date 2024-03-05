package piscine

func Sqrt(nb int) int {
	if nb%2 != 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		nb = i * i
	}
	return nb
}
