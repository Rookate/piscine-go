package piscine

func LastRune(s string) rune {
	tmp := []rune(s)
	return tmp[len(tmp)-1]
}
