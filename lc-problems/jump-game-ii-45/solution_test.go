package jump

import (
	"testing"
)

func TestJumpCorrectness(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        int
	}{
		{
			description: "success",
			input:       []int{2, 3, 1, 1, 4},
			want:        2,
		},
		{
			description: "empty",
			input:       []int{0},
			want:        0,
		},
		{
			description: "ones",
			input:       []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:        11,
		},
		{
			description: "zero_somewhere",
			input:       []int{1, 1, 1, 1, 2, 0, 1, 1, 1, 1, 1, 1},
			want:        10,
		},
		{
			description: "bounds_check",
			input:       []int{2, 1},
			want:        1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			res := jump(tc.input)
			if res != tc.want {
				t.Errorf("res != want: %d != %d", res, tc.want)
			}
		})
	}
}
