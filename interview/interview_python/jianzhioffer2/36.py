# 二叉搜索树转换成双向链表

from binarytree import Node

# 打印，直接用binarytree打印会出错
def Print(t):
    while t != None:
        print(t.value, end='->')
        t = t.right
    print()

# 思路：中序遍历，缓存下上一个节点的指针
def convert(t, preNode):
    if t == None:
        return None

    convert(t.left, preNode)
    # 调整指针，左指针指向上一个节点，上一个节点的右指针指向当前节点
    t.left = preNode
    if preNode != None:
        preNode.right = t
    print(t.value)
    
    #缓存节点
    preNode = t
    convert(t.right, preNode)

def ConvertTree2List(t):
    if t == None:
        return None

    preNode = None
    convert(t, preNode)

    #t 在中间，升序需要从左边去查找
    cur = t.left
    while cur != None:
        cur = cur.left
    return cur

t = Node(10)
t.left = Node(6)
t.right =  Node(14)
t.left.left = Node(4)
t.left.right = Node(8)
t.right.left = Node(12)
t.right.right = Node(16)

#print(t)
re = ConvertTree2List(t)
Print(re)
