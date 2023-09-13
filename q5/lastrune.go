package piscine

func LastRune(s string) rune {
	_c := 0
	for index := range s {
		_c = index
	}
	x := []rune(s)
	return x[_c]
}
