package piscine

func ActiveBits(n int) uint {
	compte := 0
	for n > 1 {
		compte += n & 1
		n >>= 1
	}
	return uint(compte)
}
