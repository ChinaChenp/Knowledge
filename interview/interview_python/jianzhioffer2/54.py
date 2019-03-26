# 二叉搜索树k大的数

from binarytree import tree, Node

# 二叉搜索树中序遍历就是递增序列
def FindKthNode(t, k):
    if t == None or k <= 0:
        return None
    return FindKthNodeCore(t, k)

def FindKthNodeCore(t, k):
    target = None
    if t.left != None:
        target = FindKthNode(t.left, k)

    print('------------', t.value)
    if target == None:
        #print(t.value)
        if k == 1:
            target = t
        k -= 1

    if target == None and t.right != None:
        target = FindKthNode(t.right, k)
    return target

t = tree(3)
print(t)
print(FindKthNode(t, 3))