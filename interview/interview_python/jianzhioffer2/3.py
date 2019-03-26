# 数据中重复的数字

# 在一个长度为n的数组里的所有数字都在0 ~ n-1的范围内。数组中某些数字是重复的，
# 但是不知道有几个数字重复了，也不知道数字重复了几次。请找出数组中任意重复的数字。
# 例如，如果输入长度为7的数字{2， 3，1，0，2， 5， 3}，那么对应输出是重复的数字2或者3

#思路：0 ~ n-1个数字没有重复，排序后数字i将出现在i的位置
def duplicate(arr):
    if arr == None:
        return None

    for i in range(0, len(arr)):
        #如果相等则不需要改变
        if arr[i] == i:
            continue
        else:
            #不相等时看下是否和目标位置相等，相等即找到了，否则就交换
            index = arr[i]
            if arr[index] == arr[i]:
                print(arr[index])
            else:
                arr[index], arr[i] = arr[i], arr[index]
arr = [2, 3, 1, 0, 2, 5, 3]
duplicate(arr)
