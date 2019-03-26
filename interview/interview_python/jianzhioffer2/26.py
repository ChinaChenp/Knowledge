# 树的子结构(剑指offer26)
# 思路先找到子树根节点，然后在对比两个树是否相等

from binarytree import Node, tree

# 匹配子树结构,t1是主树，t2是子树
def matchSubTree(t1, t2):
    if t2 == None:
        return True
    if t1 == None:
        return False
    if t1.value != t2.value:
        return False
    return matchSubTree(t1.left, t2.left) and matchSubTree(t1.right, t2.right)

# 采用前序遍历先访问根节点的方式
def IsSubTree(t1, t2):
    if t1 == None or t2 == None:
        return False
    
    # 找到一个根节点相同的，则继续判断
    result = False
    if t1.value == t2.value:
        result = matchSubTree(t1, t2)
    
    # 没有找到根节点相同的继续查找
    if result == False:
        result = IsSubTree(t1.left, t2)
    if result == False:
        result = IsSubTree(t1.right, t2)
    return result

t = tree()
subTree = Node(5)
subTree.left = Node(1)
subTree.left.right = Node(12)
print('----------------', t, subTree)
print('子树：', IsSubTree(t, subTree))
