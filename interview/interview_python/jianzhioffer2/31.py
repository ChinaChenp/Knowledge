# 栈的压入、弹出序列

# 思路：依次压入in序列直到遇到out的第一个值
# 1) 如果一直没遇到返回false
# 2）如果遇到弹出栈顶，此时如果栈顶和下一个元素相等则继续
# 3）否则继续把in压入队列
# 4）如果in已经压入完毕，但是栈顶不等于out第一个元素则返回false
def IsPopOrder(input, out):
    if input == None or out == None:
        return False

    in_index = out_index = 0

    stack = []
    while in_index < len(input):
        # len(stack) == 0第一次压入, 一直压入知道栈顶和弹出序列相等
        while in_index < len(input) and (len(stack) == 0 
        or stack[-1] != out[out_index]):
            stack.append(input[in_index])
            in_index += 1

         # 查看栈顶是否和出栈值相等
        if stack[-1] == out[out_index]:
             re = stack.pop()
             print('re', re)
             out_index += 1
        else:
            return False
    return True

l1 = [1, 2, 3, 4, 5]
l2 = [4, 5, 3, 2, 1]
print(IsPopOrder(l1, l2))
