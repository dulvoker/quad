package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')
		for stroka := 0; stroka < x-2; stroka++ {
			z01.PrintRune('B')
		}
		if x > 1 {
			z01.PrintRune('C')
		}
		if x > 0 {
			z01.PrintRune('\n')
		}
	}
	if y > 1 && x > 0 {
		for stolbec := 0; stolbec < y-2; stolbec++ {
			z01.PrintRune('B')
			for stroka := 0; stroka < x-2; stroka++ {
				z01.PrintRune(' ')
			}
			if x > 1 {
				z01.PrintRune('B')
			}
			if x > 0 {
				z01.PrintRune('\n')
			}
		}
	}
	if y > 1 && x > 0 {
		z01.PrintRune('A')
		for stroka := 0; stroka < x-2; stroka++ {
			z01.PrintRune('B')
		}
		if x > 1 {
			z01.PrintRune('C')
		}
		if x > 0 {
			z01.PrintRune('\n')
		}
	}
}
