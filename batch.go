package slice_helper

import "fmt"

func Batch[T any](slice []T, batchSize int) ([][]T, error) {
	if batchSize < 1 {
		return nil, fmt.Errorf("batch.Batch: batchSize must be greater than 0")
	}
	result := make([][]T, 0, batchSize)

	for i := 0; i < len(slice); i += batchSize {
		j := i + batchSize
		if j > len(slice) {
			j = len(slice)
		}

		result = append(result, slice[i:j])
	}

	return result, nil
}
