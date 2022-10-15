def height(root):
    # One line solution
    # return -1 if not root else max(height(root.left) + 1, height(root.right) + 1)
    
    if root == None:
        return -1
    left = height(root.left)
    right = height(root.right)
    
    if left > right:
        return left + 1
    return right + 1