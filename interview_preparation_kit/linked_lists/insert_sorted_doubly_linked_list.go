package linkedlists

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	newNode := DoublyLinkedListNode{}
	newNode.data = data

	if llist == nil {
		return &newNode
	}

	if data <= llist.data {
		newNode.next = llist
		llist.prev = &newNode
		return &newNode
	}

	rest := sortedInsert(llist.next, data)
	llist.next = rest
	rest.prev = llist
	return llist
}
