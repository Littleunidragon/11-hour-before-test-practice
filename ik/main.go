package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	val         int
	left, right *node
}

func ik(r io.Reader) (int, int) {
	var n, m int
	fmt.Fscan(r, &n, &m)
	var comp []int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &comp[i])
	}
	var t *node
	var minComp, days int
	ik := 1
	for _, i := range comp {
		if height(t) < m {
			insert(t, comp[i])
		}
		if height(t) == m {
			days++
			for t.left != nil {
				minComp = t.val
				t = t.left
				if minComp <= ik {
					ik++
				}
				deleteNode(t, minComp)
			}
		}
	}
	return days, ik
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return root
	}
	if k < root.val {
		root.left = deleteNode(root.left, k)
	} else if k > root.val {
		root.right = deleteNode(root.right, k)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			root.val = min(root.right)
			root.right = deleteNode(root.right, root.val)
		}
	}
	return root
}

func height(root *node) int {
	if root == nil || (root.left == nil && root.right == nil) {
		return 0
	}

	if height(root.left) > height(root.right) {
		return height(root.left) + 1
	} else {
		return height(root.right) + 1
	}
}

func min(root *node) int {
	for root.left != nil {
		root = root.left
	}
	return root.val
}

func insert(t *node, a int) *node {
	if t.val < a {
		insert(t.right, a)
	}
	if t.val > a {
		insert(t.left, a)
	}
	return t
}

func main() {
	r := bufio.NewReader(os.Stdout)
	ik(r)
}
