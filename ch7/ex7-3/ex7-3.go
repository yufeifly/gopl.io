package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	res := ""
	if t == nil {
		return res
	}
	res += t.left.String()
	res = fmt.Sprintf("%s %d", res, t.value)
	res += t.right.String()
	return res
}

func buildTree(data []int) *tree {
	var root = new(tree)
	for _, v := range data {
		root = add(root, v)
	}
	return root
}

func add(t *tree, e int) *tree {
	if t == nil {
		t = new(tree)
		t.value = e
		return t
	}

	if e < t.value {
		t.left = add(t.left, e)
	} else {
		t.right = add(t.right, e)
	}
	return t
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	root := buildTree(data)
	fmt.Println(root)

	//空指针
	fmt.Println(new(tree))

	//只有根节点
	root = new(tree)
	root.value = 100
	fmt.Println(root)

	//没有右子树
	data = []int{5, 4, 3, 2, 1}
	root = buildTree(data)
	fmt.Println(root)

	//没有左子树
	data = []int{1, 3, 2, 4, 5}
	root = buildTree(data)
	fmt.Println(root)
}