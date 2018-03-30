package main

import (
	"fmt"
	"learngo/tree"

	"golang.org/x/tools/container/intsets"
)

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}
func main() {
	var root tree.Node
	root = tree.Node{
		Value: 3,
	}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()
	root.Traverse()
	fmt.Println()
	fmt.Println(nodes)
	fmt.Println(root)
	fmt.Println("======")
	testSparse()
	root.Traverse2()
	fmt.Println("----")
	count := 0
	root.TraverseFunc(func(n *tree.Node) {
		count++
	})

	fmt.Println("count=", count)
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)
}
