package piscine

func Map(f func(int) bool, a []int) []bool {
	mappedSlice := make([]bool, len(slice))
	for i, v := range slice {
		mappedSlice[i] = f(v)
	}
	return mappedSlice
}
