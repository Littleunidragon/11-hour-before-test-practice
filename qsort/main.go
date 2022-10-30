package main

import "fmt"

func partition(s []int, l, r int) (int, int) {
	p := s[r/2]
	for l <= r {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l <= r {
			s[l], s[r] = s[r], s[l]
			l, r = l+1, r-1
		}
	}
	return l, r
}

func smallest(s []int, n int) int {
	l, r := 0, len(s)-1
	n = n - 1
	for l < r {
		i, j := partition(s, l, r)
		if n <= j {
			r = j
		} else if n >= i {
			l = i
		} else {
			break
		}
	}
	return s[n]
}
func main() {
	n := 2
	s := []int{9, 7, 6, 2, 5, 0, -3, 5, 4, 2, 32, 7, 74, 83, 24, 72, 6, 5, 2, 1, 0, -13, -2, 3, 5, 6, 0, 128, 132, 0, -24, 52, 10, 200, 62, 262, 907}
	fmt.Println(smallest(s, n))
	fmt.Println(s)
}
