#矩阵相乘

# 根据wikipedia上的介绍：两个矩阵的乘法仅当第一个矩阵A的行数和另一个矩阵B的列数相等时才能定义。如A是m×n矩阵，
# B是n×p矩阵，它们的乘积AB是一个m×p矩阵，它的一个元素其中 1 ≤ i ≤ m, 1 ≤ j ≤ p。

#假设是3X2矩阵
def MulMatrix(arr_a, arr_b):
    arr_c = [[0, 0],[0, 0]]
    for i in range(0, 2):
        for j in range(0, 2):
            #arr_c[i][j] = 0
            for k in range(0, 3):
                arr_c[i][j] += arr_a[i][k] * arr_b[k][j]
    return arr_c

arr_a = [[1, 2, 3], [4, 5, 6]]
arr_b = [[1, 4], [2, 5], [3, 6]]
print(MulMatrix(arr_a, arr_b))