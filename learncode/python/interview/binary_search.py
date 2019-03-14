
#二分查找算法
def binary_search(list, key):
    low, high = 0, len(list) - 1

    while low < high:
        mid = (low + high) // 2
        guess = list[mid]

        if guess == key:
            return True
        elif guess > key:
            high = mid - 1
        else:
            low = mid + 1
    return False

list = [1, 3, 6, 7, 9, 10, 12, 15]

re = binary_search(list, 3)
print(re)

re = binary_search(list, 16)
print(re)