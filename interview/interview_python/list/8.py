# 题目 ： 按某个给定值将链表划分为左边小于这个值，右边大于这个值的新链表 
# 如一个链表 为 1 -> 4 -> 5 -> 2 给定一个数 3 则划分后的链表为 1-> 2 -> 4 -> 5
# 思路：新建立2个链表，然后把大于key的插入到右链表，小于key的插入到左链表，最后把两个链表拼接在一起即可

import list_common as List

def PartitionList(head, key):
    if head == None:
        return None

    left_list, right_list = None, None #左右链表的有游标
    left_head, right_head = None, None #左右链表的头结点

    cur = head
    while cur != None:
        if cur.value <= key:
            if left_list == None:
                left_head = left_list = cur
            else:
                left_list.next = cur
                left_list = cur
        else:
            if right_list == None:
                right_head = right_list = cur
            else:
                right_list.next = cur
                right_list = cur 
        cur = cur.next 

    #拼接左右链表
    right_list.next = None
    left_list.next = right_head 
    return left_head

l2 = List.Node(6)
l2.next = List.Node(1)
l2.next.next = List.Node(4)
l2.next.next.next = List.Node(3)
l2.next.next.next.next = List.Node(5)
l2.next.next.next.next.next = List.Node(1)

re=PartitionList(l2, 3)
List.Print(re)

