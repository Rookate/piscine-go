package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := false

	for _, char := range s {
		if char == '-' && result == 0 {
			negative = true
		} else if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	if negative {
		result = -result
	}
	return result
}
