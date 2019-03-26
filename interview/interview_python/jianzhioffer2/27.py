# 二叉树的镜像

from binarytree import tree, build

def Mirror(t):
    if t == None:
        return
    #if t.left == None and t.right == None:
    #    return
    t.left, t.right = t.right, t.left
    Mirror(t.left)
    Mirror(t.right)

print('--------------')
t = build([1, 3, 8, 2, 4, 6, 0, 9])
print(t)
Mirror(t)
print(t)