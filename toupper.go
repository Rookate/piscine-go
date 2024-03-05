package piscine

func ToUpper(s string) string {
	boh := []rune(s)
	for i := 0; i < len(boh); i++ {
		if boh[i] >= 'a' && boh[i] <= 'z' {
			boh[i] = boh[i] - 32
		}
	}
	return string(boh)
}
