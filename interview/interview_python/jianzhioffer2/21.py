# 调整数组中奇数位于偶数前面
def Reorder(arr):
    if arr == None:
        return None

    beg, end = 0, len(arr) - 1
    while beg < end:
        # 从左边找第一个偶数
        while beg < end and arr[beg] % 2 != 0:
            beg = beg + 1
        # 从右边找第一个奇数
        while beg < end and arr[end] % 2 == 0:
            end = end - 1
        
        if beg < end:
            arr[beg], arr[end] = arr[end], arr[beg]
    return arr

arr = [3, 2, 9, 2, 0, 1, 5, 7, 2, 20]
print(Reorder(arr))