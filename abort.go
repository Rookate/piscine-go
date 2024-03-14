package piscine

func Abort(a, b, c, d, e int) int {
	tmp := []int{a, b, c, d, e}
	for i := 0; i < len(tmp)-1; i++ {
		for j := 0; j < len(tmp)-i-1; j++ {
			if tmp[i] < tmp[i+1] {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
	}
	return tmp[2]
}
