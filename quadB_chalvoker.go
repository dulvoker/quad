package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	for vert := 0; vert < y; vert++ {
		if vert == 0 || vert == y-1 {
			for horiz := 0; horiz < x; horiz++ {
				if horiz == 0 || horiz == x-1 {
					if horiz == 0 && vert == 0 || vert == y-1 && horiz != 0 {
						z01.PrintRune('/')
					} else {
						z01.PrintRune(92)
					}
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
		} else {
			for horiz2 := 0; horiz2 < x; horiz2++ {
				if horiz2 == 0 || horiz2 == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
