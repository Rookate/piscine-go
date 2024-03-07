package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, ch := range args {
		result += ch
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
