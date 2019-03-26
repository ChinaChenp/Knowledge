# 数组中的逆序对
# 可以考虑归并排序，不过N2算法已经很快了

def InversePairs(arr):
    for i in range(0, len(arr)-1):
        for j in range(i+1, len(arr)):
            if arr[i] > arr[j]:
                print(arr[i], arr[j])

arr = [7, 5, 6, 4]
InversePairs(arr)