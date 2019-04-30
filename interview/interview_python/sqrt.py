# 求根号x的值，保留N为小数

# 思路：二分逼近法,结束条件是前后两值只差小于最小要求精度

def Sqrt(n, digit):
    if n == 0:
        return 
    
    # 位数精度
    precision = 1
    for i in range(0, digit):
        precision = precision / 10

    beg, end = 1, n
    while abs(beg - end) > precision: #beg和end差值小于最小精度差就不可能继续划分 
        mid = (beg + end) / 2
        if mid*mid > n:
            end = mid
        else:
            beg = mid
    return beg, end

print(Sqrt(3, 3))