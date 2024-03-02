package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if i == 1 && j == 1 {
				z01.PrintRune('o')
				continue

			}
			if i == x && j == y {
				z01.PrintRune('o')
				continue
			}
			if i == x && j == 1 {
				z01.PrintRune('o')
				continue
			}
			if i == 1 && j == y {
				z01.PrintRune('o')
				continue
			}
			if j == 1 || j == y {
				z01.PrintRune('-')
				continue
			}
			if i == 1 || i == x {
				z01.PrintRune('|')
				continue
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}

}
