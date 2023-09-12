package piscine

func LastRune(s string) rune {
	_c := 0
	for index := range s {
		_c = index
	}
	_arr := []rune(s)
	return _arr[_c]
}
