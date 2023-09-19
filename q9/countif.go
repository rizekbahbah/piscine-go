package piscine

func CountIf(f func(string) bool, tab []string) int {
	computer := 0
	for _, v := range tab {
		if f(v) {
			computer++
		}
	}
	return computer
}
