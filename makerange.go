package piscine

func MakeRange(min, max int) []int {
	ln := max - min
	var bbb []int

	if min >= max {
		return nil
	}

	bbb = make([]int, ln)

	for i := 0; i < ln; i++ {
		bbb[i] = min
		min++
	}
	return bbb
}
