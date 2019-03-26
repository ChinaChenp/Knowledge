# 选择数组的最小数字(默认递增,没有重复数字)

# 思路：因为是默认递增的那么最小的值一定在数组中间
# 则取中间值：
# 中间值 > 头部，相当于中间值还没有包括最小值 beg 调整到中间值
# 中间值 < 头部（另一半数组肯定是增序）, end 调整到中间值
# 如果end - beg = 1表示两个指针刚好在临界点取最小值即可

def SortMin(arr):
    beg, end = 0, len(arr) - 1

    while beg < end:
        mid = (beg + end) // 2
        if end - beg == 1:
            return arr[end] #min(arr[beg], arr[end])
        
        if arr[mid] > arr[beg]:
            beg = mid
        elif arr[mid] < arr[end]:
            end = mid

arr = [2, 3, 4, 5, 1]
print(SortMin(arr))

arr = [2, 3, 4, 5, 1]
print(SortMin(arr))

arr = [5, 1, 2, 3, 4]
print(SortMin(arr))

