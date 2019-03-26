# 最长不包含重复字符的子串长度

def MaxSubStrLenght(arr):
    if arr == None:
        return 0
    
    map_index = [0]*256 #存储字符上一次对应的位置
    max_len, last_index = 0, 0
    for i in range(0, len(arr)):
        tmp_index = map_index[ord(arr[i]) - ord('0')]
        if tmp_index > last_index:
            max_len = max(max_len, i - tmp_index)
            # 更新最新游标为上一次下一个值
            last_index = tmp_index + 1 

        # 存储最新字符最新的位置
        map_index[ord(arr[i]) - ord('0')] = i
    max_len = max(max_len, len(arr) - last_index)
    return max_len

arr = ['a', 'b', 'c', 'a', 'b', 'c', 'b', 'b']
#arr = ['1', '1', '1']
print(MaxSubStrLenght(arr))
