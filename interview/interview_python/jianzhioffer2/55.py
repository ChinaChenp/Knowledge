# 二叉树的深度

from binarytree import Node, tree

def GetDepth(t):
    if t == None:
        return 0 
    left = GetDepth(t.left)
    right = GetDepth(t.right)

    return max(left, right) + 1

t = tree(3)
print(t, GetDepth(t))