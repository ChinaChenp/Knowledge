# 求无序数组的中位数

import heapq

def get_mid_num(l):
    #先把前二分之一的数据插入堆中
    length = (len(l) + 1) / 2
    hq = []
    for i in range(0, length):
        heapq.heappush(hq, l[i])
    
    # 剩余部分逐个遍历，如果比堆顶小则替换，否则就舍弃
    top = heapq.