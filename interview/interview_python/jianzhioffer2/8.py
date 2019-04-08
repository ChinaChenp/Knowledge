# 二叉树中序遍历下一个节点

# 思路：
# 1)如果二叉树有右子树，则下一个节点就是右子树中最左边的节点
# 2)如果没有右子树，则需要向上找到他父节点的左节点的节点

# 待调试 无法遍历 parent 不存在
def InViewNextNode(t, key):
    if t == None or key == None:
        return 
    
    # 如果右子树不为空
    next_node = None
    if t.right != None:
        cur = t.right
        while cur.left != None:
            cur = cur.left
        next_node = cur
    
    # 向上查找第一个父节点的左节点
    elif t.parent != None:
        cur, parent_node = t, t.parent 
        while parent_node != None and parent_node.right != t:
            cur = parent_node
            parent_node = parent_node.next
        next_node = parent_node
    return next_node