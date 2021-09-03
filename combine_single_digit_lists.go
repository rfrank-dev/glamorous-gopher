package main

/**
 * What's better than learning a language with a LinkedList?
 */
func CombineSingleDigitLists(leftList, rightList *IntegerNode) *IntegerNode {
	return combineSingleDigitSubLists(leftList, rightList, 0)
}

func combineSingleDigitSubLists(leftList, rightList *IntegerNode, carry int) *IntegerNode {
	if leftList == nil && rightList == nil {
		// Terminal case
		if carry == 0 {
			// Terminate the list
			return nil
		} else {
			// If there is a carry, we include it in a new terminal node.
			return &IntegerNode{nil, carry}
		}

	} else if leftList == nil {
		// Complete one list, just append the other
		return combineSingleDigitSubLists(&IntegerNode{}, rightList, carry)
	} else if rightList == nil {
		// Complete one list, just append the other
		return combineSingleDigitSubLists(leftList, &IntegerNode{}, carry)
	} else {
		// Continue recursion
		sumValue := (leftList.value + rightList.value + carry)
		node := &IntegerNode{nil, sumValue % 10}
		nextCarry := sumValue / 10
		node.nextNode = combineSingleDigitSubLists(leftList.nextNode, rightList.nextNode, nextCarry)

		return node
	}
}
