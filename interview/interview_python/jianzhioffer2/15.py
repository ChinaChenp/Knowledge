# 二进制中1的个数

def NumberOf1(n):
    max_count = 0
    count, bit = 0, 1
    while bit:
        if n & bit:
            count += 1
        bit = bit << 2

        #python 跟其他语言不一样不会溢出 默认是int64
        max_count = max_count + 1
        if max_count == 31:
            break

    return count

print(NumberOf1(4))
print(NumberOf1(5))
print(NumberOf1(7))