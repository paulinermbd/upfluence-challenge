package service

import (
	"errors"
	"reflect"
	"testing"
)

func TestComputePercentiles(t *testing.T) {
	t.Run("should return error when slice is empty", func(t *testing.T) {
		data := []int{}

		result, err := ComputePercentiles(data)

		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
		if !errors.Is(err, ErrEmptySlice) {
			t.Errorf("expected ErrEmptySlice, got %v", err)
		}
	})

	t.Run("should return error when slice is nil", func(t *testing.T) {
		var data []int

		result, err := ComputePercentiles(data)

		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
		if !errors.Is(err, ErrEmptySlice) {
			t.Errorf("expected ErrEmptySlice, got %v", err)
		}
	})

	t.Run("should compute percentiles for a single element", func(t *testing.T) {
		data := []int{42}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 42 {
			t.Errorf("expected p50=42, got %d", result[50.0])
		}
		if result[90.0] != 42 {
			t.Errorf("expected p90=42, got %d", result[90.0])
		}
		if result[99.0] != 42 {
			t.Errorf("expected p99=42, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles for two elements", func(t *testing.T) {
		data := []int{10, 20}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 10 {
			t.Errorf("expected p50=10, got %d", result[50.0])
		}
		if result[90.0] != 20 {
			t.Errorf("expected p90=20, got %d", result[90.0])
		}
		if result[99.0] != 20 {
			t.Errorf("expected p99=20, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles for unsorted data", func(t *testing.T) {
		data := []int{100, 50, 75, 25, 90}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 75 {
			t.Errorf("expected p50=75, got %d", result[50.0])
		}
		if result[90.0] != 100 {
			t.Errorf("expected p90=100, got %d", result[90.0])
		}
		if result[99.0] != 100 {
			t.Errorf("expected p99=100, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles for already sorted data", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 5 {
			t.Errorf("expected p50=5, got %d", result[50.0])
		}
		if result[90.0] != 9 {
			t.Errorf("expected p90=9, got %d", result[90.0])
		}
		if result[99.0] != 10 {
			t.Errorf("expected p99=10, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles for large dataset", func(t *testing.T) {
		data := make([]int, 100)
		for i := 0; i < 100; i++ {
			data[i] = i + 1
		}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 50 {
			t.Errorf("expected p50=50, got %d", result[50.0])
		}
		if result[90.0] != 90 {
			t.Errorf("expected p90=90, got %d", result[90.0])
		}
		if result[99.0] != 99 {
			t.Errorf("expected p99=99, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles with duplicate values", func(t *testing.T) {
		data := []int{5, 5, 5, 5, 5}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 5 {
			t.Errorf("expected p50=5, got %d", result[50.0])
		}
		if result[90.0] != 5 {
			t.Errorf("expected p90=5, got %d", result[90.0])
		}
		if result[99.0] != 5 {
			t.Errorf("expected p99=5, got %d", result[99.0])
		}
	})

	t.Run("should not modify original slice", func(t *testing.T) {
		data := []int{100, 50, 75, 25, 90}
		original := make([]int, len(data))
		copy(original, data)

		_, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(data, original) {
			t.Errorf("original slice was modified: expected %v, got %v", original, data)
		}
	})

	t.Run("should handle data with zero values", func(t *testing.T) {
		data := []int{0, 0, 0, 1, 2}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 0 {
			t.Errorf("expected p50=0, got %d", result[50.0])
		}
		if result[90.0] != 2 {
			t.Errorf("expected p90=2, got %d", result[90.0])
		}
		if result[99.0] != 2 {
			t.Errorf("expected p99=2, got %d", result[99.0])
		}
	})

	t.Run("should compute percentiles for exactly 3 elements", func(t *testing.T) {
		data := []int{1, 2, 3}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 2 {
			t.Errorf("expected p50=2, got %d", result[50.0])
		}
		if result[90.0] != 3 {
			t.Errorf("expected p90=3, got %d", result[90.0])
		}
		if result[99.0] != 3 {
			t.Errorf("expected p99=3, got %d", result[99.0])
		}
	})

	t.Run("should return all three percentiles", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != 3 {
			t.Errorf("expected 3 percentiles, got %d", len(result))
		}
		if _, ok := result[50.0]; !ok {
			t.Error("missing p50")
		}
		if _, ok := result[90.0]; !ok {
			t.Error("missing p90")
		}
		if _, ok := result[99.0]; !ok {
			t.Error("missing p99")
		}
	})

	t.Run("should handle very large numbers", func(t *testing.T) {
		data := []int{1000000, 2000000, 3000000, 4000000, 5000000}

		result, err := ComputePercentiles(data)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result[50.0] != 3000000 {
			t.Errorf("expected p50=3000000, got %d", result[50.0])
		}
		if result[90.0] != 5000000 {
			t.Errorf("expected p90=5000000, got %d", result[90.0])
		}
		if result[99.0] != 5000000 {
			t.Errorf("expected p99=5000000, got %d", result[99.0])
		}
	})
}

func TestComputePercentiles_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		data        []int
		wantP50     int
		wantP90     int
		wantP99     int
		wantErr     error
		shouldError bool
	}{
		{
			name:        "empty slice",
			data:        []int{},
			wantErr:     ErrEmptySlice,
			shouldError: true,
		},
		{
			name:        "single element",
			data:        []int{42},
			wantP50:     42,
			wantP90:     42,
			wantP99:     42,
			shouldError: false,
		},
		{
			name:        "two elements",
			data:        []int{10, 20},
			wantP50:     10,
			wantP90:     20,
			wantP99:     20,
			shouldError: false,
		},
		{
			name:        "sorted ascending",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantP50:     5,
			wantP90:     9,
			wantP99:     10,
			shouldError: false,
		},
		{
			name:        "unsorted data",
			data:        []int{100, 50, 75, 25, 90},
			wantP50:     75,
			wantP90:     100,
			wantP99:     100,
			shouldError: false,
		},
		{
			name:        "all duplicates",
			data:        []int{5, 5, 5, 5, 5},
			wantP50:     5,
			wantP90:     5,
			wantP99:     5,
			shouldError: false,
		},
		{
			name:        "with negative values",
			data:        []int{-10, -5, 0, 5, 10},
			wantP50:     0,
			wantP90:     10,
			wantP99:     10,
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ComputePercentiles(tt.data)

			if tt.shouldError {
				if err == nil {
					t.Error("expected an error but got none")
				}
				if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result[50.0] != tt.wantP50 {
				t.Errorf("p50: expected %d, got %d", tt.wantP50, result[50.0])
			}
			if result[90.0] != tt.wantP90 {
				t.Errorf("p90: expected %d, got %d", tt.wantP90, result[90.0])
			}
			if result[99.0] != tt.wantP99 {
				t.Errorf("p99: expected %d, got %d", tt.wantP99, result[99.0])
			}
		})
	}
}
