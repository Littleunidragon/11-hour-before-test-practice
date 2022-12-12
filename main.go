package main

import (
	"fmt"
	"io"
)

func IsHeap(s []int) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < (len(s)/2)-1; i++ {
		if (2*i)+1 >= s[i] || (2*i)-1 >= len(s)-1 {
			if s[(i-1)/2] > s[i] {
				return false
			}
		}
	}
	return true
}

func swords(r io.Reader) []int {
	var age []int
	var n, k int
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &age[i])
	}
	fmt.Println(n, k, age)
	return []int{0}
}

func main() {
	var s, a, d []int
	s = append(s, 1, 5, 3, 8, 7, 4, 9, 13, 15)
	a = append(a, 78, 5, 3, 8, 7, 4, 9, 13, 15)
	d = append(d, 1, 3, 2, 5, 8, 9, 6, 6, 7, 9, 3)
	fmt.Println(s)
	fmt.Println(IsHeap(s))
	fmt.Println(a)
	fmt.Println(IsHeap(a))
	fmt.Println(d)
	fmt.Println(IsHeap(d))
	//r := bufio.NewReader(os.Stdout)
	//swords(r)
}
