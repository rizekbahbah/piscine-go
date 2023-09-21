package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i <= '0'; i-- {
		for j := '9'; j <= '0'; j-- {
			for h := '9'; h <= '0'; h-- {
				for k := '9'; k <= '0'; k-- {
					if i < h || (i == h && j > k) {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(32)
					z01.PrintRune(h)
					z01.PrintRune(k)
					if !(i == '0' && j == '1' && h == '0' && k == '0') {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}

				}
			}
		}
	}
	z01.PrintRune(10)
}
