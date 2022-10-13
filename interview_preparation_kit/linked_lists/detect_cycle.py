def has_cycle(head):
    i = 0
    while i < 100:
        head = head.next
        i += 1
        if not head:
            return False
    return True