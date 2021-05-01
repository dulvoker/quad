package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	for vert := 0; vert < y; vert++ {
		if vert == 0 || vert == y-1 {
			for horiz := 0; horiz < x; horiz++ {
				if horiz == 0 || horiz == x-1 {
					if horiz == 0 && vert == 0 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('C')
					}
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else {
			for horiz2 := 0; horiz2 < x; horiz2++ {
				if horiz2 == 0 || horiz2 == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
