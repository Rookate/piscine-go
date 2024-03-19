package main

import "strconv"

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	if from <= to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i < to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i > to {
				result += ", "
			}
		}
	}
	result += "\n"
	return result
}

func main() {
	print(FromTo(1, 10))
	print(FromTo(10, 1))
	print(FromTo(10, 10))
	print(FromTo(100, 10))
}
