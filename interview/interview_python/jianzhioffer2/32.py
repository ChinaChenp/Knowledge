# 层次遍历二叉树

from binarytree import tree

def TreeLevelView(t):
    if t == None:
        return None
    qu = []
    qu.append(t)

    while len(qu) != 0:
        t = qu.pop(0)
        print(t.value, end = ', ')

        if t.left != None:
            qu.append(t.left)
        if t.right != None:
            qu.append(t.right)
t = tree()
print(t)
TreeLevelView(t)