#给定一个字符串，求它的最长回文子串的长度。
#最容易想到的办法是枚举所有的子串，分别判断其是否为回文。这个思路初看起来是正确的，但却做了很多无用功，如果一个长的子串包含另一个短一些的子串，那么对子串的回文判断其实是不需要的。
#枚举每个中心字段向两边展开判断是否是回文字符串


def IsPalindrome(arr, center):
    Min = min(center, len(arr) - center - 1)
    beg, end = center - Min, center + Min

    length = 0
    while beg < end:
        length = length + 1
        if arr[beg] != arr[end]:
            return length
        beg = beg + 1
        end = end - 1
    return length

def LongestPalindrome(arr):
    if len(arr) < 2:
        return 0, 1
    max_index, length = 0, 0
    for v in range(1, len(arr)):
        re = IsPalindrome(arr, v)
        if length <= re:
            max_index, length = v, re
    return max_index, length
arr = list("23uf0a0fdedf0abdedfgfdedba")
index, length = LongestPalindrome(arr)
print(index, length)