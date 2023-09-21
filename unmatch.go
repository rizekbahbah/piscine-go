package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		x := 0
		for _, el := range a {
			if el == res {
				x++
			}
		}
		if x == 1 || x%2 == 1 {
			return res
		}
	}
	return -1
}
