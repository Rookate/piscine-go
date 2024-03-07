package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	result := -1
	for i, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if len(s[result+1:i]) > 0 {
				if s[result+1:i] != " " {
					tab = append(tab, s[result+1:i])
				}
				result = i
			}
		}
	}
	tab = append(tab, s[result+1:])
	return tab
}
