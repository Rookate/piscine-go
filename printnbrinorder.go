package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var array []int
	eachValue := 0
	// Conversion de l'entier en un tableau d'entiers
	for n != 0 {
		eachValue = n % 10
		n /= 10
		array = append(array, eachValue)
	}
	// Tri du tableau dans l'ordre croissant
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	// Impression des chiffres dans l'ordre croissant
	for i := 0; i < len(array); i++ {
		z01.PrintRune(rune(array[i] + '0'))
	}
}
