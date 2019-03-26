# 二叉树中和为某一值的路径

from binarytree import tree, Node

# 注意路劲定义只从根节点到叶子节点
def FindPath(t, need_sum):
    if t == None:
        return 

    path = [] #节点缓存
    cur_sum = 0
    FindAllPath(t, need_sum, cur_sum, path)

def FindAllPath(t, need_sum, cur_sum, path):
    # 把当前节点压入栈，并且总和累加
    path.append(t.value)
    cur_sum += t.value

    # 查看是否只叶子节点并且累加和等于目标sum
    isLeaf = (t.left == None and t.right == None)
    if isLeaf and cur_sum == need_sum:
        #打印路劲
        for v in path:
            print(v, end=', ')
        print()
    if t.left != None:
        FindAllPath(t.left, need_sum, cur_sum, path)
    if t.right != None:
        FindAllPath(t.right, need_sum, cur_sum, path)
    # 如果是叶子结点又不满足要求，退回到父结点，删除当前节点
    path.pop()
    
t = Node(10)
t.left = Node(5)
t.right = Node(12)
t.left.left = Node(4)
t.left.right = Node(7)

print(t) 
FindPath(t, 22)