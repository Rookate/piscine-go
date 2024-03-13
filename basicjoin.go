package piscine

func BasicJoin(elems []string) string {
	resultat := ""
	for _, ch := range elems {
		resultat += ch
	}
	return resultat
}
