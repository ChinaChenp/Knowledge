#字符串移动

#字符串为*号和26个字母的任意组合，把*号都移动到最左侧，把字母移到最右侧并保持相对顺序不变，
# 要求时间和空间复杂度最小

#解法 从后往前找
def ShiftChar(arr):
    slow = fast = len(arr) - 1
    
    while True:
        #slow 从后面找第一个*字符
        for i in range(fast, -1, -1):
            if arr[i] == '*':
                slow = i
                break

        #从slow开始找王前找第一个字符
        for i in range(slow, -1, -1):
            if arr[i] >= 'A' and arr[i] <= 'z':
                fast = i
                break

        print(arr)
        if fast < slow:
            arr[fast], arr[slow] = arr[slow], arr[fast]
            slow = slow - 1
            fast = slow
        else: # fast == slow 两个
            break
    return arr

arr = list("jf*da*ji*eo*fjac*")
print(ShiftChar(arr))
