# 和为s的连续正数序列
# 注意条件：1）连续正数从1开始 2）是增续
# 思路：快慢指针，cur > sum时舍弃最小值，cur < sum 追加最大值

def CalcSum(beg, end):
    total = 0
    while beg <= end:
        total += beg
        beg += 1
    return total

def FindSeqNum(sum):
    if sum < 3:
        return None

    slow, fast = 1, 2
    end = (sum + 1) // 2 #结束条件
    while slow < end:
        cur = CalcSum(slow, fast)
        #print('------', slow, fast, cur)
        if cur == sum:
            print(slow, fast)
            slow += 1
        elif cur > sum:
            slow += 1
        else:
            fast += 1
    return None


FindSeqNum(15)
    
