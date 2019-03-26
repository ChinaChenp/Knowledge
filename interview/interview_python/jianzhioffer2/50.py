# 第一个只出现一次的字符

# 思路：用hash[256]
def FistNotRepeatChar(arr):
    if arr == None:
        return None

    # 先hash所有字符, 统计出现次数
    hash_map = [0]*256
    for v in arr:
        hash_map[ord(v)] += 1
    #print(hash_map)

    for v in arr:
        if hash_map[ord(v)] == 1:
            return v
    
    return arr[0]

arr = list("xabaccdyeffxbd")
print(FistNotRepeatChar(arr))