# 合并两个无序数组合并，有可能有重复
# 比如：l1 = {1, 3, 5, 7, 9, 10, 22}, l2 = {0, 2, 3, 7, 9, 25}
# 输出 {3, 7, 9}

# 求交集，有重复
def ArrIntersection(l1, l2):
    if l1 == None or l2 == None:
        return []
    
    i = j = 0
    out = []
    while i < len(l1) and j < len(l2):
        if l1[i] == l2[j]:
            out.append(l1[i])
            i += 1
            j += 1
        elif l1[i] > l2[j]:
            j += 1
        else:
            i += 1
    return out

# 求并集合
def ArrUnion(l1, l2):
    if l1 == None or l2 == None:
        return []
    
    i = j = 0
    out = []
    while i < len(l1) and j < len(l2):
        if l1[i] == l2[j]:
            i += 1
            j += 1
        elif l1[i] > l2[j]:
            out.append(l2[j])
            j += 1
        else:
            out.append(l1[i])
            i += 1
    return out


l1 = [1, 3, 5, 7, 9, 9, 10, 22]
l2 = [0, 2, 3, 7, 9, 9, 25]

re = ArrIntersection(l1, l2)
print(re)

re = ArrUnion(l1, l2)
print(re)