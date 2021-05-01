package piscine

import "github.com/01-edu/z01"

func Quada(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('o')
		for a := 1; a <= x-2; a++ {
			z01.PrintRune('-')
		}
		if x > 1 {
			z01.PrintRune('o')
		}
		if x > 0 {
			z01.PrintRune('\n')
		}
	}
	if y > 1 && x > 0 {
		for b := 1; b <= y-2; b++ {
			z01.PrintRune('|')
			for a := 1; a <= x-2; a++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('|')
			}
			if x > 0 {
				z01.PrintRune('\n')
			}
		}

	}
	if y > 1 && x > 0 {
		z01.PrintRune('o')
		for a := 1; a <= x-2; a++ {
			z01.PrintRune('-')
		}
		if x > 1 {
			z01.PrintRune('o')
		}
		z01.PrintRune('\n')
	}
}
