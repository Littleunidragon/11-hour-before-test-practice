package main

import "testing"

func TestSmallest(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		n    int
		want int
	}{
		{[]int{1, 2, 3, 4}, 1, 1},
		{[]int{67, 23, 45, 799, 0}, 3, 45},
	} {
		if smallest(tc.s, tc.n) != tc.want {
			t.Errorf("QuickSort(%v, %v) = %v, want = %v", tc.s, tc.n, smallest(tc.s, tc.n), tc.want)
		}
	}
}
