package piscine

func AppendRange(min, max int) []int {
	var bbb []int
	for i := min; i < max; i++ {
		bbb = append(bbb, i)
	}
	return bbb
}
