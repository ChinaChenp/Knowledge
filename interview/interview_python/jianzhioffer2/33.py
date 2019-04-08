# 二叉搜索树的后续遍历序列
# 输入一个序列判断该序列是不是二叉树的后续遍历结果

def VerifySquBst(l):
    if l == None or len(l) <= 0:
        return False
    if len(l) == 1: #叶子节点
        return True

    # 根据root找到到左序列和右序列
    root = l[-1]
    print('root=', root, l)

    # 从左往后找第一个大于root的节点(搜索树左边比root小，右边比root大)
    index = 0
    for i in range(0, len(l) - 1): #最后个节点是跟节点无需判断
        if l[i] > root:
            index = i
            break
    
    # 判断右边是否都比root大，否则异常
    for j in range(index, len(l) - 1):
        if l[j] < root: 
            print('============', l[j], root)
            return False
    
    # 递归左边和右边
    left, right = False, False
    if index > 0:
        left = VerifySquBst(l[:index])
    if index < len(l) - 1:
        right = VerifySquBst(l[index:len(l)-1])
    return left and right

arr = [5, 7, 6, 9, 11, 10, 8]
print(VerifySquBst(arr))
