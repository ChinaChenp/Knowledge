# 打印从1到最大的n位数

# 方式很简单，但是要考虑越界的问题，因而采用字符串做加法

def IncOne(arr):
    length = len(arr) - 1
    take_over = 0
    for i in range(length, -1, -1):
        now = int(arr[i]) - int('0') + take_over

        #print('in', arr, now)
        # 如果是个位需要+1
        if length == i:
            now += 1
        
        # 如果需要进位
        if now >= 10:
            arr[i] = str(now - 10 + int('0'))
            take_over = 1

            # 如果是最高位还需要进位
            if i == 0:
                return ['1'] + arr
        else:
            arr[i] = str(now + int('0'))
            take_over = 0
    return arr

re = [0]
for i in range(0, 101):
    re = IncOne(re)
    print(''.join(re))
