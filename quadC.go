package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if i == 1 && j == 1 {
				z01.PrintRune('A')
				continue

			}
			if i == x && j == 1 {
				z01.PrintRune('A')
				continue
			}
			if i == x && j == y {
				z01.PrintRune('C')
				continue
			}
			if i == 1 && j == y {
				z01.PrintRune('C')
				continue
			}
			if j == 1 || j == y {
				z01.PrintRune('B')
				continue
			}
			if i == 1 || i == x {
				z01.PrintRune('B')
				continue
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}
