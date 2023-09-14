package piscine

func ConcatParams(args []string) string {
	Str := ""
	ln := 0
	for i := range args {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		if i != ln {
			Str = Str + args[i] + "\n"
		} else {
			Str = Str + args[i]
		}
	}
	return Str
}
