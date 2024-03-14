package piscine

func Compact(ptr *[]string) int {
	var tmp []string
	for _, ch := range *ptr {
		if ch != "" {
			tmp = append(tmp, ch)
		}
	}
	*ptr = tmp
	return len(tmp)
}
