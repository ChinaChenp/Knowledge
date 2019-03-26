#行列递增矩阵的查找

#首先直接定位到最右上角的元素，再配以二分查找，比要找的数（6）大就往左走，比要找数（6）的小就往下走，直到找到要找的数字（6）为止，这个方法的时间复杂度O（m+n）
def YoungMatrix(arr, key, row, col):
    i, j = 0, col - 1

    while True:
        cur = arr[i][j]
        if cur > key:
            j -= 1
        elif cur < key:
            i += 1
        else:
            return arr[i][j]
        if j < 0 or i >= col:
            return False

arr = [[1, 2, 8, 9],
       [2, 4, 9, 12],
       [4, 7, 10, 13],
       [6, 8, 11, 15],
      ]
print(YoungMatrix(arr, 6, 4, 4))