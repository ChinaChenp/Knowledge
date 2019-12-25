#coding=utf-8
#链表反转

import list_common as List

#图解方式  https://blog.csdn.net/xungjhj/article/details/76401184
def list_reverse(head):
    if head == None or head.next == None:
        return head
    
    p, q = head, head.next 
    head.next = None
    while q: # 注意从第二个值开始循环
        tmp = q.next
        q.next = p
        p = q
        q = tmp
    head = p  # 最后p已经指向none了
    return head

head = List.Node(1)
head.next = List.Node(2)
head.next.next = List.Node(3)
head.next.next.next = List.Node(4)

List.Print(head)
new_list = list_reverse(head)
List.Print(new_list)