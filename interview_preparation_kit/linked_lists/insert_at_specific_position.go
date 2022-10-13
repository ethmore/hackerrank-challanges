package linkedlists

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	newNode := SinglyLinkedListNode{}
	newNode.data = data

	if position < 0 {
		return llist
	}
	if position == 0 {
		llist = &newNode
		return llist
	}

	var prev *SinglyLinkedListNode
	ptr := llist

	for i := int32(0); i < position; i++ {
		if i == position-1 {
			prev = ptr
		}
		ptr = ptr.next
	}

	newNode.next = ptr
	prev.next = &newNode

	return llist
}
