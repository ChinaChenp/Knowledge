# 局部最大（小）值
# 给定无序数组arr，已知arr中任意两个相邻的数都不相等。写一个函数，只需返回arr中任意一个局部最小出现的位置即可

# 解法 
# 1)arr长度为1时，arr[0]是局部最小。 
# 2)arr的长度为N(N>1)时， 
# ①如果arr[0]＜arr[1]，那么arr[0]是局部最小； 
# ②如果arr[N-1]＜arr[N-2]，那么arr[N-1]是局部最小； 
# ③如果0＜i＜N-1，既有arr[i]＜arr[i-1]，又有arr[i]＜arr[i+1]，那么arr[i]是局部最小

#考虑找最小值
def getLessIndex(arr):
    if len(arr) == 0:
        return -1
    
    #考虑边界左右两边情况
    if len(arr) == 1:
        return 0
    if len(arr) == 2 and arr[0] < arr[1]:
        return 0
    if arr[len(arr) - 1] < arr[len(arr) - 2]:
        return len(arr) - 1
    
    #二分查找, 边界不用考虑了
    beg, end = 1, len(arr) - 2
    while beg <= end:
        mid = (beg + end) // 2
        #如果左边有递减趋势则在左边找
        if arr[mid] > arr[mid - 1]:
            end = mid - 1
        elif arr[mid] > arr[mid + 1]:
            beg = mid + 1
        else:
            return mid
    return -1

arr = [2, 9, 10, 4, 2, 1, 11, 14, 8, 25]
print(getLessIndex(arr))