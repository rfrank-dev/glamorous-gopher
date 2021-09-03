package main

import (
	"testing"
)

func TestAppendItem(t *testing.T) {
	listNode := IntegerNode{nil, 5}
	listNode.append(2)
	listNode.append(6)
	listNode.append(10)
}
