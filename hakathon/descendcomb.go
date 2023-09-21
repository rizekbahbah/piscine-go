package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for x := '9'; x >= '0'; x-- {
		for y := '9'; y >= '0'; y-- {
			for z := '9'; z >= '0'; z-- {
				for k := '9'; k >= '0'; k-- {
					if (x > z) || (x == z && y > k) {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(k)
						if !(x == '0' && y == '1' && z == '0' && k == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
