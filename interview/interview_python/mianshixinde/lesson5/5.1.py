#最大连续乘积子串

def MaxProductSubstring1(arr):
    max_value = arr[0]
    for i in range(0, len(arr)):
        tmp = arr[i]
        for j in range(i + 1, len(arr)):
            tmp *= arr[j]
            if tmp > max_value:
                max_value = tmp
    return max_value

arr = [-2.5, 4, 0, 3, 0.5, 8, -1]
print(MaxProductSubstring1(arr))

# 解法二(因为有正负数)
# 考虑到乘积子序列中有正有负也还可能有0，我们可以把问题简化成这样：
# 数组中找一个子序列，使得它的乘积最大；同时找一个子序列，使得它的乘积最小（负数的情况）。
# 因为虽然我们只要一个最大积，但由于负数的存在，我们同时找这两个乘积做起来反而方便。也就是说，
# 不但记录最大乘积，也要记录最小乘积。

def MaxProductSubstring2(arr):
    max_value = arr[0]
    min_value = arr[0]
    ret = arr[0]
    for i in range(1, len(arr)):
        #同时有正负数，最小值肯定为负数，负数随时可以翻身成最大值，如果只有正数就简单了
        max1, min1 = max_value * arr[i], min_value * arr[i]
        max_value = max(max1, min1, arr[i])
        min_value = min(max1, min1, arr[i])

        ret = max(ret, max_value)
        print(max_value, min_value, ret)
    return ret
arr = [-2.5, 4, 0, 3, 0.5, 8, -2, -1, -40]
print(MaxProductSubstring2(arr))
arr = [-2.5, -1, 0, -1, -2, 3, -10]
print(MaxProductSubstring2(arr))
arr = [2.5, 1, 100, 0, 1, 2, 3, 10]
print(MaxProductSubstring2(arr))
