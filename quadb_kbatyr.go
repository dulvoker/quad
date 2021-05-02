package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune(47)
		for a := 0; a < x-2; a++ {
			z01.PrintRune('*')
		}
		if x > 1 {
			z01.PrintRune(92)
		}
		if x > 0 {
			z01.PrintRune('\n')
		}
	}
	if y > 1 && x > 0 {
		for b := 0; b < y-2; b++ {
			z01.PrintRune('*')
			for a := 0; a < x-2; a++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('*')
			}
			if x > 0 {
				z01.PrintRune('\n')
			}
		}

	}
	if y > 1 && x > 0 {
		z01.PrintRune(92)
		for a := 0; a < x-2; a++ {
			z01.PrintRune('*')
		}
		if x > 1 {
			z01.PrintRune(47)
		}
		z01.PrintRune('\n')
	}
}
