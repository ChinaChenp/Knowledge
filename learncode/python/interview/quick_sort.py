#快速排序
def quick_sort(arr):
    if len(arr) < 2:
        return arr

    pivot = arr[0]
    less = [i for i in arr[1:] if i <= pivot]
    greater = [j for j in arr[1:] if j > pivot]
    return quick_sort(less) + [pivot] + quick_sort(greater)

arr = [9, 3, 10, 2, 7, 5, 0, 8, -1, 4, 3]
print(quick_sort(arr))
