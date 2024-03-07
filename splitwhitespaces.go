package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	result := -1
	for i, ch := range s {
		if ch == ' ' || ch == '	' || ch == '\n' {
			if len(s[result+1:i]) > 0 {
				tab = append(tab, s[result+1:i])
			}
			result = i
		}
	}
	if result != len(s)-1 {
		tab = append(tab, s[result+1:])
	}
	return tab
}
