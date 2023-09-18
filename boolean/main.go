package piscine

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	lengthOfArg := 0

	sentence := os.Args
	for index := range sentence {
			lengthOfArg = index 
	}


func isEven(lengthOfArg) == true {
	printStr("I have an even number of arguments")
} else {
	printStr("I have an odd number of arguments")
}
}

func printStr(str string) {
	arrayStr := []rune(str)

	lengthOfArg := 0

	for ind := range arrayStr {
		lengthOfArg = ind
	}

	for i := 0; i < lengthOfArg; i++ {

		z01.PrintRune(rune(arrayStr[i]))
	}
	z01.PrintRune('\n')
}

func is Even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}




