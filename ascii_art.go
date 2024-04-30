package main

func Runify(strings []string) [][]rune {
	runes := make([][]rune, len(strings))
	for i, str := range strings {
		runes[i] = []rune(str)
	}
	return runes
}
