package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if y == 1 {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('A')
			for z := 1; z < y-1; z++ {
				if z <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
			for v := 1; v < y-1; v++ {
				z01.PrintRune('B')

				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')

			}
			z01.PrintRune('A')

			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
		}
	}
}