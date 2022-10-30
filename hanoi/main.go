package main

import "fmt"

func hanoi(n, from, to int, f func(from, to int)) {
	if n == 1 {
		f(from, to)
	} else if from == 1 && to == 3 {
		hanoi(n, from, 6-from-to, f)
		hanoi(n, 6-from-to, to, f)
	} else {
		hanoi(n-1, from, 6-from-to, f)
		hanoi(1, from, to, f)
		hanoi(n-1, 6-from-to, to, f)
	}
}

func main() {
	n := 5
	hanoi(n, 1, 3, func(from, to int) { fmt.Println(n, from, to) })
}
