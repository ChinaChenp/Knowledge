#寻找和为定值的两个数
#输入一个数组和一个数字，在数组中查找两个数，使得它们的和正好是输入的那个数字。
#要求时间复杂度是O(N)。如果有多对数字的和等于输入的数字，输出任意一对即可。
#例如输入数组1、2、4、7、11、15和数字15。由于4+11=15，因此输出4和11。

#默认升序
def two_number(arr, need):
    beg, end = 0, len(arr) - 1
    while beg < end:
        #print(beg, end, arr[beg], arr[end])
        total = arr[beg] + arr[end]
        if total == need:
            return arr[beg], arr[end]
        elif total > need:
            end = end - 1
        else:
            beg = beg + 1
    return -1, -1

arr = [1, 2, 4, 7, 11, 15]
print(two_number(arr, 15))
print(two_number(arr, 20))

