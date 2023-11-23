package batch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBatch(t *testing.T) {
	cases := []struct {
		name      string
		arr       []int64
		batchSize int
		expected  [][]int64
		wantError assert.ErrorAssertionFunc
	}{
		{
			name:      "empty_slice",
			arr:       []int64{},
			batchSize: 10,
			expected:  [][]int64{},
			wantError: assert.NoError,
		},
		{
			name:      "nil_slice",
			arr:       nil,
			batchSize: 10,
			expected:  [][]int64{},
			wantError: assert.NoError,
		},
		{
			name:      "size_1_full_filling",
			arr:       []int64{1, 2, 3, 4},
			batchSize: 1,
			expected:  [][]int64{{1}, {2}, {3}, {4}},
			wantError: assert.NoError,
		},
		{
			name:      "size_2_not_full_filling_random_sort",
			arr:       []int64{5, 1, 3},
			batchSize: 2,
			expected:  [][]int64{{5, 1}, {3}},
			wantError: assert.NoError,
		},
		{
			name:      "size_10_single_batch",
			arr:       []int64{4, 7, 9, 2, 5},
			batchSize: 10,
			expected:  [][]int64{{4, 7, 9, 2, 5}},
			wantError: assert.NoError,
		},
		{
			name:      "invalid_batch_size",
			arr:       []int64{4, 7, 9, 2, 5},
			batchSize: 0,
			expected:  nil,
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "batch.Batch: batchSize must be greater than 0")

				return true
			},
		},
		{
			name:      "invalid_batch_size_2",
			arr:       []int64{4, 7, 9, 2, 5},
			batchSize: -1,
			expected:  nil,
			wantError: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "batch.Batch: batchSize must be greater than 0")

				return true
			},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			batches, err := Batch(tt.arr, tt.batchSize)
			tt.wantError(t, err)

			assert.Equal(t, tt.expected, batches)
		})
	}
}
