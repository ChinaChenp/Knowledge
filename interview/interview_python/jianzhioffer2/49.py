# 丑数

# 穷举算法
def IsUglyNum(cur):
    while cur % 2 == 0:
        cur = cur // 2
    while cur % 3 == 0:
        cur = cur // 3
    while cur % 5 == 0:
        cur = cur // 5
    if cur == 1:
        return True
    return False

def UglyNum(num):
    cout = 0 #丑数数量
    cur = 1  #自增num
    while cout < num:
        cur += 1
        if IsUglyNum(cur):
            print(cur)
            cout += 1
    return cur

print(UglyNum(20))   
print()  

# 解法二:
# 丑数肯定是由前面的某个丑数乘以2、3、5得来的，但是乘以2、3、5的丑数位置各不相同
def UglyNum1(num):
    mul1, mul2, mul3 = 0, 0, 0
    all_ugly = [0]*num

    cout, all_ugly[0] = 1, 1 #初始化第一个丑数
    while cout < num:
        # 先计算下一个丑数，这个丑数肯定只有某个丑数构成，因而下面三个if只有一个会生效
        next_ugly = min(all_ugly[mul1]*2, all_ugly[mul2]*3, all_ugly[mul3]*5)
        if all_ugly[mul1] * 2 <= next_ugly:
            mul1 += 1
        if all_ugly[mul2] * 3 <= next_ugly:
            mul2 += 1
        if all_ugly[mul3] * 5 <= next_ugly:
            mul3 += 1
        all_ugly[cout] = next_ugly
        print(next_ugly)
        cout += 1
    return all_ugly[cout - 1]


print(UglyNum1(9))
