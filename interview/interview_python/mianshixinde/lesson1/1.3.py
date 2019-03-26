#输入一个由数字组成的字符串，把它转换成整数并输出。例如：输入字符串"123"，输出整数123。
#给定函数原型int StrToInt(const char *str) ，实现字符串转换成整数的功能，不能使用库函数atoi。

def atoi(str):
    if len(str) == 0:
        return 0

    #判断正负号,默认正整数
    flag = True
    if str[0] == '-':
        flag = False
        str = str[1:]
    elif str[0] == '+':
        flag = True
        str = str[1:]
    else:
        flag = True

    out = 0
    for v in str:
        if v < '0' or v > '9':
            break
        num = ord(v) - ord('0')
        out = out * 10 + num
    if flag :
        return out
    return out * -1

print(atoi("+3243 44"))
print(atoi("+324fff3 44"))
print(atoi("-3241fff3 44"))
