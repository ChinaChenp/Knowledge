#1.1 旋转字符串
#给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部，
# 如把字符串“abcdef”前面的2个字符'a'和'b'移动到字符串的尾部，
# 使得原字符串变成字符串“cdefab”。请写一个函数完成此功能，要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)。

def left_shift_one(s):
    tmp, l = s[0], len(s) - 1
    for i in range(0, l):
        s[i] = s[i+1]
    s[l] = tmp
    return s

def left_shift(s, n):
    if len(s) == 0:
        return ""
    n = n % len(s)

    while n > 0:
        s = left_shift_one(s)
        n = n - 1
    return s

s = list('123456789')
print("".join(left_shift(s[:], 2)))
print("".join(left_shift(s[:], 11)))