package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		if char == ' ' {
			return 0
		}
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}
	return result
}
