#字符串全排列
#输入一个字符串，打印出该字符串中字符的所有排列。
#例如输入字符串abc，则输出由字符a、b、c 所能排列出来的所有字符串
#abc、acb、bac、bca、cab 和 cba。

def Permutation(arr, index):
    if len(arr) == 0:
        return
    if len(arr) - 1 == index:
        print(''.join(arr))
    for v in range(index, len(arr)):
        arr[v], arr[index] = arr[index], arr[v]
        Permutation(arr, index + 1)
        arr[index], arr[v] = arr[v], arr[index]
str = ['1', '2', '3']
Permutation(str, 0)