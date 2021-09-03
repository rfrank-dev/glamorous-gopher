package main

import (
	"testing"
)

func TestCombineSimpleLists(t *testing.T) {
	leftNode := IntegerNode{nil, 1}
	leftNode.append(1)
	leftNode.append(1)

	rightNode := IntegerNode{nil, 1}
	rightNode.append(1)
	rightNode.append(1)
	rightNode.append(1)
	rightNode.append(1)
	rightNode.append(1)

	compareArrays(t, leftNode.to_a(), []int{1, 1, 1})
	compareArrays(t, rightNode.to_a(), []int{1, 1, 1, 1, 1, 1})

	result := CombineSingleDigitLists(&leftNode, &rightNode)
	compareArrays(t, result.to_a(), []int{2, 2, 2, 1, 1, 1})
}

func TestCombineDifferentLengthLists(t *testing.T) {
	leftNode := IntegerNode{nil, 1}
	leftNode.append(2)
	leftNode.append(3)
	leftNode.append(4)
	leftNode.append(5)
	leftNode.append(6)

	rightNode := IntegerNode{nil, 1}
	rightNode.append(2)
	rightNode.append(3)

	result := CombineSingleDigitLists(&leftNode, &rightNode)
	compareArrays(t, result.to_a(), []int{2, 4, 6, 4, 5, 6})
}

func TestCombineDifferentLengthListsWithCarry(t *testing.T) {
	leftNode := IntegerNode{nil, 9}
	leftNode.append(9)
	leftNode.append(9)
	leftNode.append(9)

	rightNode := IntegerNode{nil, 9}
	rightNode.append(9)
	rightNode.append(9)
	rightNode.append(9)
	rightNode.append(9)
	rightNode.append(9)

	compareArrays(t, leftNode.to_a(), []int{9, 9, 9, 9})
	compareArrays(t, rightNode.to_a(), []int{9, 9, 9, 9, 9, 9})

	result := CombineSingleDigitLists(&leftNode, &rightNode)
	compareArrays(t, result.to_a(), []int{8, 9, 9, 9, 0, 0, 1})
}

func compareArrays(t *testing.T, leftList, rightList []int) {
	if len(leftList) != len(rightList) {
		t.Errorf("The lists do not have equal length: %v != %v", leftList, rightList)
		t.Fail()
	}

	for idx, val := range leftList {
		if val != rightList[idx] {
			t.Errorf("The lists do not have equal contents: %v != %v", leftList, rightList)
			t.Fail()
		}
	}
}
