package entities

import (
	"bytes"
	"fmt"
)

type Node struct {
	value map[string]int
	next  *Node
}

func NewNode(value map[string]int, next *Node) *Node {
	return &Node{value: value, next: next}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(node *Node) *Node {
	n.next = node

	return n
}

func (n *Node) Start() int {
	return n.value["start"]
}

func (n *Node) End() int {
	return n.value["end"]
}

func (n *Node) String() string {
	pair := n.value
	start := pair["start"]
	end := pair["end"]

	var buffer bytes.Buffer
	buffer.WriteString("Start:")
	buffer.WriteString(fmt.Sprint(start))
	buffer.WriteString(", End:")
	buffer.WriteString(fmt.Sprint(end))

	if n.next != nil {
		buffer.WriteString("; Next: ")
		buffer.WriteString(fmt.Sprintf("%v", n.next))
	}

	return buffer.String()
}
