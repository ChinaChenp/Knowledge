#和为s的数字

def FindSumNum(arr, sum):
    if arr == None:
        return

    beg, end = 0, len(arr) - 1
    while beg < end:
        cur = arr[beg] + arr[end]

        if cur == sum:
            return arr[beg], arr[end]
        elif cur > sum:
            end -= 1
        else: #cur < sum:
            beg += 1
    return None

arr = [1, 2, 4, 7, 11, 15]
print(FindSumNum(arr, 15))
