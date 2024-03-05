package piscine

func IsAlpha(s string) bool {
	for _, ch := range s {
		if (ch < 'a' || ch > 'z') && (ch < '0' || ch > '9') && (ch < 'A' || ch > 'Z') {
			return false
		}
	}
	return true
}
