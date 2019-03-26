# 数组中唯一出现的一个数
# 数组中有一个数只出现一次，其他的都出现三次

# 思路：把每个数的每一位分别相加，那么出现三次的刚好能被三整除，为1的位置就是目标数

def FindOnceNum(arr):
    if arr == None:
        return None
    
    # 分别统计所以数字每一位1的数量
    bit_count = [0]*32
    for num in arr:
        mask = 1
        for j in range(31, -1, -1): #要反向遍历，最低位在右边
            if num & mask != 0:     #要判断 bit != 0，判断 bit == 0有问题，这种方式是嘴个判断每一位
                bit_count[j] += 1
            mask = mask << 1
    
    # 还原目标数字
    #print(bit_count)
    result = 0
    for i in range(0, 32):
        result = result << 1
        result += bit_count[i] % 3
    return result

arr = [1, 2, 2, 1, 7, 2, 1, 3, 4, 3, 3, 4, 4]
print(FindOnceNum(arr))

