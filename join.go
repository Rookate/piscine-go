package piscine

func Join(strs []string, sep string) string {
	resultat := ""
	for i, ch := range strs {
		if i > 0 {
			resultat += sep
		}
		resultat += ch
	}
	return string(resultat)
}
