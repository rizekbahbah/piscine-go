package piscine

func IsUpper(str string) bool {
	_arr := []rune(str)
	ln := 0
	for i := range _arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if _arr[i] < 'A' || _arr[i] > 'Z' {
			return false
		}
	}
	return true
}
