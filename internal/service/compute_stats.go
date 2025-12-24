package service

import (
	"errors"
	"math"
	"sort"
)

var (
	ErrEmptySlice        = errors.New("cannot calculate percentile of empty slice")
	ErrInvalidPercentile = errors.New("percentile must be between 0 and 100")
)

// ComputePercentiles compute all percentiles in one pass
func ComputePercentiles(data []int) (map[float64]int, error) {
	percentiles := []float64{50, 90, 99}
	if len(data) == 0 {
		return nil, ErrEmptySlice
	}

	for _, p := range percentiles {
		if p < 0 || p > 100 {
			return nil, ErrInvalidPercentile
		}
	}

	sorted := make([]int, len(data))
	copy(sorted, data)
	sort.Ints(sorted)

	results := make(map[float64]int, len(percentiles))
	for _, p := range percentiles {
		// Nearest rank method as I work with int data this is acceptable
		rank := math.Ceil((p / 100.0) * float64(len(sorted)))
		index := int(rank) - 1 // conversion en index (base 0)

		// Protecting index out of bounds
		if index < 0 {
			index = 0
		}
		if index >= len(sorted) {
			index = len(sorted) - 1
		}

		results[p] = sorted[index]
	}

	return results, nil
}
