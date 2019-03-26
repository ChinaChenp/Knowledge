#二分查找

def BinarySearch(arr, key):
    beg, end = 0, len(arr) - 1

    while beg <= end:
        mid = (beg + end) // 2
        value = arr[mid]
        if value > key:
            end = mid - 1
        elif value < key:
            beg = mid + 1
        else:
            return True, value
    return False, None

arr = [1, 4, 5, 6, 8, 23, 29, 30, 31]
print(BinarySearch(arr, 1))
print(BinarySearch(arr, 31))
print(BinarySearch(arr, 8))
print(BinarySearch(arr, 9))
