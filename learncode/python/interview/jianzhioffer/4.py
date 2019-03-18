# 二位数组中的查找

def Find(arr, key, row_num, col_num):
    row, col = 0, col_num - 1  # 索引位置指向右上角
    while row < row_num and col >= 0:
        v = arr[row][col]
        if v == key:
            return v
        elif v > key:
            col = col - 1
        else: # v < key
            row = row + 1
    return -1

arr = [[1, 2, 8, 9],
       [2, 4, 9, 12],
       [4, 7, 10, 13],
       [6, 8, 11, 15],
      ]
re = Find(arr, 4, 4, 4)
print(re)

re = Find(arr, 15, 4, 4)
print(re)

re = Find(arr, 8, 4, 4)
print(re)

re = Find(arr, 1, 4, 4)
print(re)

re = Find(arr, 20, 4, 4)
print(re)
