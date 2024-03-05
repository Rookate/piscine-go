package piscine

func ToLower(s string) string {
	boh := []rune(s)
	for i := 0; i < len(boh); i++ {
		if boh[i] >= 'A' && boh[i] <= 'Z' {
			boh[i] = boh[i] + 32
		}
	}
	return string(boh)
}
