package piscine

func Atoi(s string) int {
	result := 0
	p := 1
	for i, char := range s {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				p = -1
			}
		} else if char < '0' || char > '9' {
			return 0
		} else {
			result = result*10 + int(char-'0')
		}
	}
	return result * p
}
