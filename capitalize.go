package piscine

func Capitalize(s string) string {
	if len(s) <= 0 {
		return ""
	}
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if ch[i] >= 'a' && ch[i] <= 'z' || ch[i] >= 'A' && ch[i] <= 'Z' || ch[i] >= '0' && ch[i] <= '9' {
			if ch[i] >= 'a' && ch[i] <= 'z' {
				ch[i] -= 32
			}
			i++
		}
		for i < len(ch) && (ch[i] >= 'a' && ch[i] <= 'z' || ch[i] >= 'A' && ch[i] <= 'Z' || ch[i] >= '0' && ch[i] <= '9') {
			if ch[i] >= 'A' && ch[i] <= 'Z' {
				ch[i] += 32
			}
			i++
		}
	}
	return string(ch)
}
