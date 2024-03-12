package piscine

func FindNextPrime(nb int) int {
	answere := 0
	if nb < 0 {
		return -1
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			answere = i
		}
	}
	return answere
}
