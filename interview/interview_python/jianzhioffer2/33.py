# 二叉搜索树的后续遍历序列

# 待调试
def VerifySquBst(l, length):
    if l == None or length <= 0:
        return False
    
    # 根据root找到到做序列和右序列
    root = l[length - 1]

    index = 0
    for i in range(0, length-1): #length -1 最后一个节点是根节点
        if l[i] > root:
            index = i
            break
    print('--------', l[index])
    
    for j in range(index, length-1):
        if l[j] <= root: # 右边如果比root小说明异常
            print('============', l)
            return False
    
    # 递归左边和右边
    left, right = False, False
    if index > 0:
        left = VerifySquBst(l, index)
    if index < length - 1:
        right = VerifySquBst(l[index:], length - 1 - index)
    return left and right

arr = [5, 7, 6, 9, 11, 10, 8]
print(VerifySquBst(arr, len(arr)))
