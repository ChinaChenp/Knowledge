#数组中数字出现的次数

def GetFirstBit(num):
    cout = 0
    while 1 & num != 1:
        num = num >> 1
        cout += 1
    return cout

def IsBit1(num, cout):
    return num & 1<<cout

def FindNums(arr):
    # 计算所有异或总和
    total = 0
    for i in arr:
        total ^= i
    
    n = GetFirstBit(total)
    
    # 按某一位区分num
    num1, num2 = 0, 0
    for i in arr:
        if IsBit1(i, n):
            num1 ^= i
        else:
            num2 ^= i
    return num1, num2

arr = [2, 4, 3, 6, 3, 2, 5, 5]
print(FindNums(arr))
    

