package minimumswaps2

import (
	"testing"
)

func TestMinimumSwaps(t *testing.T) {
	var tests = map[string]struct {
		arr  []int32
		want int32
	}{
		"first_example": {
			arr:  []int32{7, 1, 3, 2, 4, 5, 6},
			want: 5,
		},
		"second_example": {
			arr:  []int32{4, 3, 1, 2},
			want: 3,
		},
		"third_example": {
			arr:  []int32{2, 3, 4, 1, 5},
			want: 3,
		},
		"fourth_example": {
			arr:  []int32{1, 3, 5, 2, 4, 6, 7},
			want: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := MinimumSwaps(test.arr)
			if res != test.want {
				t.Errorf("got %v, want %d", res, test.want)
			}
		})
	}
}
