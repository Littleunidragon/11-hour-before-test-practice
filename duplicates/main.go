package main

import "fmt"

type node struct {
	val         int
	right, left *node
}

func duplicates(t *node) bool {
	var tree [][]int
	var lineTree []int
	scanTree(t, &tree, 0)
	for _, i := range tree {
		for _, j := range i {
			lineTree = append(lineTree, j)
		}
	}
	for i := 1; i < len(lineTree)-1; i++ {
		if lineTree[i] == lineTree[i+1] || lineTree[0] == lineTree[len(lineTree)-1] {
			return false
		}
	}
	return true
}

func scanTree(t *node, Preresult *[][]int, level int) {
	if t == nil {
		return
	}
	if len(*Preresult) == level {
		*Preresult = append(*Preresult, []int{})
	}
	(*Preresult)[level] = append((*Preresult)[level], t.val)

	scanTree(t.left, Preresult, level+1)
	scanTree(t.right, Preresult, level+1)
}
func main() {
	t := &node{val: 1, right: &node{val: 5, right: &node{val: 9}, left: &node{val: 6}}, left: &node{val: 3}}
	fmt.Println(duplicates(t))

}
