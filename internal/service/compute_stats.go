package service

import (
	"errors"
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
		if p == 0 {
			results[p] = sorted[0]
			continue
		}
		if p == 100 {
			results[p] = sorted[len(sorted)-1]
			continue
		}

		index := (p / 100.0) * float64(len(sorted)-1)
		roundedIndex := int(index + 0.5)

		if roundedIndex >= len(sorted) {
			results[p] = sorted[len(sorted)-1]
			continue
		}

		results[p] = sorted[roundedIndex]
	}

	return results, nil
}
