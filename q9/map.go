package piscine

func Map(f func(int) bool, arr []int) []bool {
	tabRetour := []bool{}
	for _, v := range arr {
		tabRetour = append(tabRetour, f(v))
	}
	return tabRetour
}
