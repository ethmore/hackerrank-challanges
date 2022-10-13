package linkedlists

func reverse(llist *DoublyLinkedListNode) *DoublyLinkedListNode {
	if llist.next == nil {
		return llist
	}

	curr := llist
	var lastElem *DoublyLinkedListNode
	for curr != nil {
		if curr.next == nil {
			lastElem = curr
		}
		tempNextAdd := curr.next
		curr.next = curr.prev
		curr.prev = tempNextAdd

		curr = curr.prev
	}
	return lastElem
}
