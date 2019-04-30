# 在排序数字中查找数字

# 思路：用二分查找第一次出现k的位置和最后一次出现k的位置

def FindFirstK(arr, key):
    beg, end = 0, len(arr) - 1

    while beg <= end:
        mid = (beg + end) // 2

        # 相等的时候要判断左边是不是还有相对的数字
        #print(beg, end, mid, arr[beg], arr[end])
        if arr[mid] == key:
            # 左边还有key，则要在左边查找; 结束条件是mid为=0，即最左边或者下一个值不等于key
            if mid == 0 or arr[mid - 1] != key:
                return mid
            else:
                end = mid - 1
        elif arr[mid] < key:
            beg = mid + 1
        else:
            end = mid - 1
    return -1

def FindLastK(arr, key):
    beg, end = 0, len(arr) - 1

    while beg <= end:
        mid = (beg + end) // 2

        if arr[mid] == key:
            if mid == len(arr) - 1 or arr[mid + 1] != key:
                return mid
            else:
                beg = mid + 1
        elif arr[mid] > key:
            end = mid - 1
        else:
            beg = mid + 1
    return -1

def SortCommonNum(arr, key):
    left = FindFirstK(arr, key)
    if left == -1:
        return -1
    rigth = FindLastK(arr, key)
    if rigth == -1:
        return -1

    if left <= rigth:
        return rigth - left + 1
    else:
        return -1

arr = [1, 2, 3, 3, 3, 3, 4, 5]
print(SortCommonNum(arr, 3))

arr = [3, 3, 3, 3, 3, 4, 5, 5, 6]
print(SortCommonNum(arr, 3))

arr = [1, 2, 3, 3, 3, 3, 4, 5, 5, 5]
print(SortCommonNum(arr, 5))