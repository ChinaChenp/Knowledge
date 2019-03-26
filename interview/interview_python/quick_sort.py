#快速排序
def quick_sort(arr):
    if len(arr) < 2:
        return arr
    else:
        pivot = arr[0]

        more, less = [], []
        for v in range(1, len(arr)):
            if arr[v] > pivot:
                more.append(arr[v])
            else:
                less.append(arr[v])
        return quick_sort(less) + [pivot] + quick_sort(more)
        
arr = [9, 3, 10, 2, 7, 5, 0, 8, -1, 4, 3]
print(quick_sort(arr))

import heapq
#归并排序
#把两个有序数组合并
def merge(left, right):
    i, j = 0, 0  #记录剩余值
    re = []
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            re.append(left[i])
            i += 1
        else:
            re.append(right[j])
            j += 1
    #结束的时候肯定至少有一个list已经走到尾巴
    re += left[i:]
    re += right[j:]
    return re

def merge_sort(arr):
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    return merge(left, right)

arr = [9, 3, 10, 2, 7, 5, 0, 8, -1, 4, 3]
print(merge_sort(arr))
