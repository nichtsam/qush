package main

import (
	"math/rand/v2"

	"golang.org/x/exp/maps"
)

func generateUniqueRandoms[V any](options []V) []V {
	optionCount := len(options)
	resultCount := rand.IntN(optionCount)
	result := make(map[int]V)

	for len(result) < resultCount {
		randomIndex := rand.IntN(optionCount)
		_, exist := result[randomIndex]
		if exist {
			continue
		}
		result[randomIndex] = options[randomIndex]
	}

	return maps.Values(result)
}
