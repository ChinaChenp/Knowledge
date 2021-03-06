#两个有序链表合并

import list_common as List

# 思路1：递归
def MergeSortList(l1, l2):
    if l1 == None:
        return l2
    if l2 == None:
        return l1
    
    # 正向排序
    cur = None
    if l1.value < l2.value:
        cur = l1 #cur串联l1节点
        cur.next = MergeSortList(l1.next, l2) #串联下一个节点
    else:
        cur = l2
        cur.next = MergeSortList(l1, l2.next)
    return cur

# 思路2: 遍历
def MergeSortList1(l1, l2):
    if l1 == None:
        return l2
    if l2 == None:
        return l1
    
    # 取头结点
    head = cur = None
    if l1.value < l2.value:
        cur = head = l1
        l1 = l1.next
    else:
        cur = head = l2
        l2 = l2.next
    
    # 每串联一个节点当前节点cur往后移动
    while l1 != None and l2 != None:
        if l1.value < l2.value:
            cur.next = l1
            l1 = l1.next
        else:
            cur.next = l2
            l2 = l2.next
        cur = cur.next # 当前游标指针往后移动
        List.Print(head)
    
    if l1 != None:
        cur.next = l1
    if l2 != None:
        cur.next = l2
    return head

a, b, c, d, e = List.Node(1), List.Node(3), List.Node(7), List.Node(9), List.Node(10)
head1 = a
head1.next = b
head1.next.next = c
head1.next.next.next = d
head1.next.next.next.next = e

a1, b1, c1, d1, e1 = List.Node(0), List.Node(3), List.Node(3), List.Node(8), List.Node(11)
head2 = a1
head2.next = b1
head2.next.next = c1
head2.next.next.next = d1
head2.next.next.next.next = e1
head2.next.next.next.next.next = List.Node(12)

#re = MergeSortList(head1, head2)
#Print(re)

re = MergeSortList1(head1, head2)
List.Print(re)
