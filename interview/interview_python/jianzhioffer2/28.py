# 判断二叉树是否是对称的

from binarytree import Node

def isSymmetric(t):
    return isSymmetricInner(t, t)


def isSymmetricInner(t1, t2):
    if t1 == None and t2 == None:
        return True

    if t1 == None or t2 == None:
        return False
    
    if t1.value != t2.value:
        return False
        
    left = isSymmetricInner(t1.left, t2.right) # 注意这里，对称和镜像不同之处
    right = isSymmetricInner(t1.right, t2.left)
    return left and right

t = Node(1)
t.left = t.right = Node(8)
t.left.left = t.right.right = Node(9)
print(t)
print(isSymmetric(t))