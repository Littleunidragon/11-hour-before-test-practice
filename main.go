package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	val      int
	children []int
}

func ans(r io.Reader) {
	var n, m int
	fmt.Fscan(r, &n, &m)
	cats := make([]int, n)
	for i := range cats {
		fmt.Fscan(r, &cats[i])
	}
	coord := make(map[int]int)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		coord[x] = append(coord[x], x)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	ans(r)
}
