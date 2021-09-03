package main

import "fmt"

type IntegerNode struct {
	nextNode *IntegerNode
	value    int
}

func (node IntegerNode) hasNext() bool {
	return node.nextNode != nil
}

func (node *IntegerNode) append(val int) {
	if node.hasNext() {
		node.nextNode.append(val)
	} else {
		node.nextNode = &IntegerNode{nil, val}
	}
}

func (node IntegerNode) print() {
	fmt.Println(node.value)

	if node.hasNext() {
		node.nextNode.print()
	}
}

func (node IntegerNode) to_a() []int {
	if node.hasNext() {
		return append([]int{node.value}, node.nextNode.to_a()...)
	} else {
		return []int{node.value}
	}
}
