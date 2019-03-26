#选择排序

def select_sort(list):
    for i in range(0, len(list)-1):
        min = i
        for j in range(i+1, len(list)):
            if list[j] > list[min]:
                min = j
                
        if min != i:       
            list[min], list[i] = list[i], list[min]
    return list

list = [5, 3, 6, 2, 1, 0]
re = select_sort(list)
print(re)