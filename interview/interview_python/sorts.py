#快速排序
def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    povit = arr[0]

    left, right = [], []
    for i in range(1, len(arr)):
        if arr[i] < povit:
            left.append(arr[i])
        else:
            right.append(arr[i])
    return quick_sort(left) + [povit] + quick_sort(right)

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


def select_sort(list):
    for i in range(0, len(list)-1):
        min = i
        for j in range(i+1, len(list)):
            if list[j] < list[min]:
                min = j
                
        if min != i:       
            list[min], list[i] = list[i], list[min]
    return list

arr = [9, 3, 10, 2, 7, 5, 0, 8, -1, 4, 3]
print(select_sort(arr))