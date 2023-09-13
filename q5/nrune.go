package piscine

func NRune(s string, n int) rune {
	_c := 0
	for index := range s {
		_c = index
	}
	if (_c+1) >= n && n > 0 {
		_arr := []rune(s)
		return (_arr[n-1])
	} else {
		return 0
	}
}
