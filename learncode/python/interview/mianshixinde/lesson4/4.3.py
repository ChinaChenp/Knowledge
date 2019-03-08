#出现次数超过一半的数字
# 更进一步，考虑到这个问题本身的特殊性，我们可以在遍历数组的时候保存两个值：一个candidate，用来保存数组中遍历到的某个数字；一个nTimes，表示当前数字的出现次数，其中，nTimes初始化为1。当我们遍历到数组中下一个数字的时候：
# 如果下一个数字与之前candidate保存的数字相同，则nTimes加1；
# 如果下一个数字与之前candidate保存的数字不同，则nTimes减1；
# 每当出现次数nTimes变为0后，用candidate保存下一个数字，并把nTimes重新设为1。 直到遍历完数组中的所有数字为止。
# 举个例子，假定数组为{0, 1, 2, 1, 1}，按照上述思路执行的步骤如下：
# 1.开始时，candidate保存数字0，nTimes初始化为1；
# 2.然后遍历到数字1，与数字0不同，则nTimes减1变为0；
# 3.因为nTimes变为了0，故candidate保存下一个遍历到的数字2，且nTimes被重新设为1；
# 4.继续遍历到第4个数字1，与之前candidate保存的数字2不同，故nTimes减1变为0；
# 5.因nTimes再次被变为了0，故我们让candidate保存下一个遍历到的数字1，且nTimes被重新设为1。最后返回的就是最后一次把nTimes设为1的数字1。

def FindOneNumber(arr):
    if len(arr) == 0:
        return
    times, candidate = 0, 0
    for v in arr:
        if times == 0:
            times = 1
            candidate = v
        else:
            if v == candidate:
                times = times + 1
            else:
                times = 0
    return candidate

arr = [0, 1, 2, 1, 1]
print(FindOneNumber(arr))
arr = [0, 1, 2, 1, 1, 2, 1, 2, 3, 2, 2]
print(FindOneNumber(arr))
arr = [0, 1, 2, 1, 1, 3, 3, 3, 3]
print(FindOneNumber(arr))