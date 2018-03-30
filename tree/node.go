package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
func (node *Node) SetValue(value int) {
	node.Value = value
}
func Print(node Node) {
	fmt.Println(node.Value)
}
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
