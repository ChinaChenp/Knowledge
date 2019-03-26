# 替换空格

# 实现一个函数，把字符串中的每个空格替换成"%20".例如，输入"we are happy",则输出“we%20are%20happy”

#思路：从后往前赋值, 但是要先统计空格数量预留足够长的空间（默认空间够可以从最后一位往前）
def ReplaceBlank(arr):
    # 默认空间够
    length = len(arr) - 1
    slow = fast = length

    # 从后面找第一个非空字符
    for i in range(length, -1, -1):
        if arr[i] != ' ':
            fast = i
            break
    
    # 从末尾依次拷贝
    for _ in range(fast, -1, -1):
        #替换成 %20
        if arr[fast] == ' ':
            arr[slow] = '0'
            slow = slow - 1
            arr[slow] = '2'
            slow = slow - 1
            arr[slow] = '%'
        else: #直接拷贝
            arr[slow] = arr[fast]

        #调整有标
        slow = slow - 1
        fast = fast - 1
    arr = arr[slow + 1:]
    return arr
arr = list("we are happy       ")
arr = ReplaceBlank(arr)
print(''.join(arr))
