#二叉树
#pip install binarytree

from binarytree import tree, build, Node, bst

#t = build([2, 5, 8, 3, 1, 9, 10, 13, 6, 7, 34, 19, 23, 45, 31, 24])
t = build([2, 5, 8, 3, 1, 9, 10, 13, 14, 12])
print(t)

# 求二叉树节点个数
def GetNodeNum(t):
    if t == None:
        return 0
    
    left = GetNodeNum(t.left)
    right= GetNodeNum(t.right)
    return left + right + 1

re = GetNodeNum(t)
print('节点个数:{}'.format(re))

# 求二叉树深度
def GetDepth(t):
    if t == None:
        return 0
    left = GetDepth(t.left)
    right = GetDepth(t.right)
    return max(left, right) + 1

re = GetDepth(t)
print('深度:{}'.format(re))

# 二叉树最小深度(从根节点到叶子节点最短距离)
def GetMinDepth(t):
    if t == None:
        return 0
    left = GetMinDepth(t.left)
    right = GetMinDepth(t.right)
    return min(left, right) + 1
re = GetMinDepth(t)
print('最小深度:{}'.format(re))


#前序遍历，中序遍历，后序遍历
def PreOrderTraverse(t):
    if t == None:
        return
    print(t.value, end=', ')
    PreOrderTraverse(t.left)
    PreOrderTraverse(t.right)

print('前序遍历: ')
PreOrderTraverse(t)
print()

#前序遍历非递归写法
def PreOrderTraverse1(t):
    if t == None:
        return

    stack = []
    cur = t
    while cur != None or len(stack) != 0:
        while cur != None:
            print(cur.value, end=', ')
            stack.append(cur)
            cur = cur.left
        
        if len(stack) != 0:
            cur = stack.pop()
            cur = cur.right

PreOrderTraverse1(t)
print()

def InOrderTraverse(t):
    if t == None:
        return
    InOrderTraverse(t.left)
    print(t.value, end=', ')
    InOrderTraverse(t.right)


print('中序遍历: ')
InOrderTraverse(t)
print()


def PostOrderTraverse(t):
    if t == None:
        return
    PostOrderTraverse(t.left)
    PostOrderTraverse(t.right)
    print(t.value, end=', ')


print('后序遍历: ')
PostOrderTraverse(t)
print()

#分层遍历二叉树(从上到下)
def LevelTraverse(t):
    if t == None:
        return
    l = list()
    l.append(t)

    while len(l) != 0:
        node = l.pop(0)
        print(node.value, end=', ')
        if node.left != None:
            l.append(node.left)
        if node.right != None:
            l.append(node.right)
        #print('------------', l)
    return
print('分层遍历(从上到下):')
LevelTraverse(t)
print()

#分层遍历二叉树(从下到上，从左到右)
def LevelTraverse1(t):
    if t == None:
        return
    l = list()
    l.append(t)

    out = [] #所有层list
    path = [] #每一层路劲
    cur_count, next_count = 1, 0
    while len(l) != 0:
        node = l.pop(0)
        path.append(node.value)
        cur_count -= 1
        if node.left != None:
            l.append(node.left)
            next_count += 1
        if node.right != None:
            l.append(node.right)
            next_count += 1
        if cur_count == 0: #当前层最后一个元素
            out.append(path)
            path = []
            cur_count = next_count
            next_count = 0
    out.reverse()
    return out
print('分层遍历(从下到上):')
re = LevelTraverse1(t)
print(re)

#求二叉树第K层的节点个数
def GetNodeNumKthLevel(t, k):
    if t == None or k == 0:
        return 0
    if k == 1:
        return 1
    left = GetNodeNumKthLevel(t.left, k - 1)
    right = GetNodeNumKthLevel(t.right, k - 1)
    #print('-----------', left, right)
    return left + right
re = GetNodeNumKthLevel(t, 3)
print('第k（3）层数据个数：{}'.format(re))

#叶子节点的个数
def GetLeafNodeNum(t):
    if t == None:
        return 0
    if t.left == None and t.right == None:
        return 1
    left = GetLeafNodeNum(t.left)
    right = GetLeafNodeNum(t.right)
    return left + right


re = GetLeafNodeNum(t)
print('叶子节点的个数：{}'.format(re))

# 判断两棵二叉树是否结构相同
def StructureCmp(t1, t2):
    if t1 == None and t2 == None:
        return True
    if t1 == None or t2 == None:
        return False
    
    l1 = StructureCmp(t1.left, t2.left)
    l2 = StructureCmp(t1.right, t2.right)
    return l1 and l2


re = StructureCmp(t, t)
print(re)
t2 = build([2, 5, 8, 13, 14, 10, 1, 3, 45, 39, 11, 15, 18, 19, 23, 20])
re = StructureCmp(t, t2)
print(re)

#判断一棵二叉树是不是平衡二叉树
#（1）如果二叉树为空，返回真
#（2）如果二叉树不为空，如果左子树和右子树都是AVL树并且左子树和右子树高度相差不大于1，返回真，其他返回假
def IsAVL(t):
    if t == None:
        return True

    left = GetDepth(t.left)
    right = GetDepth(t.right)
    # 深度差1不满足
    if abs(left - right) > 1:
        return False
    #在递归左边和右边
    return IsAVL(t.left) and IsAVL(t.right)

re = IsAVL(t)
print(re)

re = IsAVL(t2)
print(t2, re)

# 二叉树的镜像
def Mirror(t):
    if t == None:
        return None
    t.left, t.right = t.right, t.left
    Mirror(t.left)
    Mirror(t.right)


print('二叉树镜像:', t)
Mirror(t)
print(t)

# 二叉树的右视图,即打印每层最后一个节点
def LeftView(t):
    if t == None:
        return
    l = list()
    l.append(t)

    count = 0 #当前层处理的个数
    cur_level, next_level = 1, 0  #记录当前层和下一层个数
    while len(l) != 0:
        node = l.pop(0)
        if node.left != None:
            l.append(node.left)
            next_level += 1
        if node.right != None:
            l.append(node.right)
            next_level += 1
        if count == 0:
            print(node.value, end=', ')
        count += 1
        if count == cur_level:
            #清零计数器 
            cur_level = next_level
            count = next_level = 0
    return

# 递归的方式
# 因为每层只打印一个节点，我们可以使用一个记录每层节点的队列，同时把层数作为一个参数，当打印到某层时，
# 如果层数刚好等于队列的长度，就说明该节点就是该层遇到的第一个节点，将它加入队列并向子树递归。
def LeftView1(t, level, slice):
    #print(level, slice)
    if t == None:
        return None
    if level == len(slice):
        slice.append(t.value)
    LeftView1(t.left, level + 1, slice)
    LeftView1(t.right, level + 1, slice)

print('右视图遍历:')
LeftView(t)
print()
re = []
LeftView1(t, 0, re)
print(re)
print()

# 二叉树节点间最大距离

# 二叉树中节点的最大距离必定是两个叶子节点的距离。求某个子树的节点的最大距离，有三种情况：
# 1.两个叶子节点都出现在左子树；
# 2.两个叶子节点都出现在右子树；
# 3.一个叶子节点在左子树，一个叶子节点在右子树。只要求得三种情况的最大值，结果就是这个子树的节点的最大距离。

# int find_max_len(Node * root)
# case 1: 两个叶子节点都出现在左子树。find_max_len(root -> pLeft)
# case 2: 两个叶子节点都出现在右子树。find_max_len(root -> pRight)
# case 3: 一个出现在左子树，一个出现在右子树。distance(root -> pLeft) + distance(root -> pRight) + 2
# 其中，distance()计算子树中最远的叶子节点与根节点的距离，其实就是左子树的高度减1。

def MaxHigh(t):
    if t == None:
        return 0
    return GetDepth(t) #高度等于深度减1

def MaxDistance(t):
    if t == None:
        return 0
    if t.left == None and t.right == None:
        return 0
    
    #case1 两个叶子节点都在左子树
    lmax = MaxDistance(t.left)

    #case2 两个叶子节点都在右子树
    rmax = MaxDistance(t.right)

    #case3 一个在根节点左边、一个在右边
    lhigh = MaxHigh(t.left)
    rhigh = MaxHigh(t.right)
    #print(lmax, rmax, lhigh, rhigh)

    return max(lmax, rmax, lhigh + rhigh)

re = MaxDistance(t)
print(t)
print('最大节点距离:{}'.format(re))
print()

# 最低公共 两个节点的最低公共祖先

# 如果是二叉搜索树
# 对于二叉搜索树来说，公共祖先的值一定大于等于较小的节点，小于等于较大的节点。换言之，
# 在遍历树的时候，如果当前结点大于两个节点，则结果在当前结点的左子树里，如果当前结点小于两个节点，、
# 则结果在当前节点的右子树里
def lowestCommonAncestorOfBST(t, n1, n2):
    if t == None:
        return None
    if t.value > n1 and t.value > n2:
        return lowestCommonAncestorOfBST(t.left, n1, n2)
    if t.value < n1 and t.value < n2:
        return lowestCommonAncestorOfBST(t.right, n1, n2)
    return t


# 如果是普通二叉树
# 我们可以用深度优先搜索，从叶子节点向上，标记子树中出现目标节点的情况。如果子树中有目标节点，
# 标记为那个目标节点，如果没有，标记为null。显然，如果左子树、右子树都有标记，说明就已经找到最小公共祖先了。
# 如果在根节点为p的左右子树中找p、q的公共祖先，则必定是p本身。
# 换个角度，可以这么想：如果一个节点左子树有两个目标节点中的一个，右子树没有，
# 那这个节点肯定不是最小公共祖先。如果一个节点右子树有两个目标节点中的一个，左子树没有，
# 那这个节点肯定也不是最小公共祖先。只有一个节点正好左子树有，右子树也有的时候，才是最小公共祖先。

# 类似于后续遍历的方式
def lowestCommonAncestorOfBST1(t, n1, n2):
    # 发现目标节点则通过返回值标记该子树发现了某个目标结点
    if t == None or t.value == n1 or t.value == n2:
        return t
    # 查看左子树中是否有目标结点，没有为null
    left = lowestCommonAncestorOfBST1(t.left, n1, n2)

    # 查看右子树是否有目标节点，没有为null
    right = lowestCommonAncestorOfBST1(t.right, n1, n2)

    # 都不为空，说明做右子树都有目标结点，则公共祖先就是本身
    if left != None and right != None:
        return t
    
    # 没找到继续往上找
    if left != None:
        return left
    if right != None:
        return right
    

re = lowestCommonAncestorOfBST1(t, 8, 1)
print('最低公共祖先:{}'.format(re.value))

# 树的子结构(剑指offer26)
# 思路先找到子树根节点，然后在对比两个树是否相等

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

subTree = Node(5)
subTree.left = Node(1)
subTree.left.right = Node(12)
print('----------------', t, subTree)
print('子树：', IsSubTree(t, subTree))
