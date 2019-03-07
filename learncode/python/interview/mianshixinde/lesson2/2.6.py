#奇偶排序
#输入一个整数数组，调整数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。要求时间复杂度为O(n)。

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
        if beg < end:
            arr[beg], arr[end] = arr[end], arr[beg]
    return arr

arr = [4, 5, 2, 7, 9, 1, 4, 0, 13]
print(OddEvenSort(arr))
        