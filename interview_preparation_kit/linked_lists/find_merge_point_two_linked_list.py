def findMergeNode(head1, head2):
    currA = head1
    currB = head2
    while currA.next != None:
        temp = currA
        currA = currA.next
        temp.next = None
        
    while currB.next != None:
        currB = currB.next
        
    return currB.data