package piscine

func AlphaCount(str string) int {
	_str := []byte(str)
	_counter := 0
	for _, letter := range _str {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			_counter++
		}
	}
	return _counter
}
