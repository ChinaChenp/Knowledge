#奇偶排序
#输入一个整数数组，调整数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。要求时间复杂度为O(n)。

#相对位置会发生变化
def IsOddNumber(num):
    if num % 2 == 0:
        return True
    return False

def OddEvenSort(arr):
    beg, end = 0, len(arr) - 1
    while beg < end:
        #从左边寻找第一个偶数
        while beg < end and not IsOddNumber(arr[beg]):
            beg = beg + 1
        #从右边寻找第一个奇数
        while beg < end and IsOddNumber(arr[end]):
            end = end - 1
        arr[beg], arr[end] = arr[end], arr[beg]
    return arr

arr = [4, 5, 2, 7, 9, 1, 4, 0, 13]
print(OddEvenSort(arr))


#输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，
# 所有的偶数位于位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

#方式一 冒泡排序思想
def OddEvenSort1(arr):
    for i in range(0, len(arr)):
        for j in range(i+1, len(arr)):
            #奇数往前面冒
            if IsOddNumber(arr[i]) and IsOddNumber(arr[j]) == False:
                arr[i], arr[j] = arr[j], arr[i]
                break
    return arr

arr = [4, 5, 2, 7, 9, 1, 4, 0, 13, 12]
print(OddEvenSort1(arr))

#方式二  插入排序，每次选一个第一个奇数
def OddEvenSort2(arr):
    for i in range(0, len(arr)):
        #每次选第一个奇数进行交换
        for j in range(i+1, len(arr)):
            if IsOddNumber(arr[j]) == False:
                arr[i], arr[j] = arr[j], arr[i]
                break
    return arr

arr = [4, 5, 2, 7, 9, 1, 4, 0, 13, 12]
print(OddEvenSort2(arr))
