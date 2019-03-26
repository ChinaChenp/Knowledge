#把数组排成最小的数
import functools

def str_cmp(x, y):
    s1 = x + y
    s2 = y + x
    
    if s1 < s2:
        return -1
    elif s1 > s2:
        return 1
    else:
        return 0

arr = ['3', '32', '321']
#print(sorted(arr, key=functools.cmp_to_key(str_cmp)))
#print(sorted(arr, key=functools.cmp_to_key(str_cmp)))
re = arr.sort(key=functools.cmp_to_key(str_cmp))
print(re, arr)
