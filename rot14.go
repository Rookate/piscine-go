package piscine

func Rot14(s string) string {
	tmp := []rune(s)
	for ch := 0; ch < len(tmp); ch++ {
		if tmp[ch] >= 'A' && tmp[ch] <= 'z' {
			tmp[ch] += 14
			if (tmp[ch] > 'Z' && tmp[ch] <= 'Z'+14) || tmp[ch] > 'z' {
				tmp[ch] -= 26
			}
		}
	}
	return string(tmp)
}
