package piscine

func ActiveBits(n int) uint {
	compte := 1
	for n > 0 {
		compte += n & 1
		n >>= 1
	}
	return uint(compte)
}
