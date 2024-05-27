package main

import (
	"math/rand/v2"

	"golang.org/x/exp/maps"
)

func generateUniqueRandoms[V any](options []V, minAmount, maxAmount int) []V {
	optionCount := len(options)
	resultCount := max(min(rand.IntN(optionCount), maxAmount), minAmount)
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
