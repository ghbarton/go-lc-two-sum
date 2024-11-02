package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		exp    []int
	}{
		{
			name:   "Test 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			exp:    []int{0, 1},
		},
		{
			name:   "Test 2",
			nums:   []int{3, 2, 4},
			target: 6,
			exp:    []int{1, 2},
		},
		{
			name:   "Test 3",
			nums:   []int{3, 3},
			target: 6,
			exp:    []int{0, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.nums, tc.target)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
