# 单链表排序

import list_common as List

# 冒泡排序 默认升序
def BubbleSort(head):
    if head == None:
        return head

    cur = head
    while cur != None:
        next_node = cur.next 

        while next_node != None and next_node.next != None:
            if next_node.value > next_node.next.value:
                next_node.value, next_node.next.value = next_node.next.value, next_node.value
            next_node = next_node.next 
        cur = cur.next 
    return head 

l2 = List.Node(1)
l2.next = List.Node(2)
l2.next.next = List.Node(6)
l2.next.next.next = List.Node(3)
l2.next.next.next.next = List.Node(4)
l2.next.next.next.next.next = List.Node(4)
l2.next.next.next.next.next.next = List.Node(9)
l2.next.next.next.next.next.next.next = List.Node(5)
l2.next.next.next.next.next.next.next.next = List.Node(5)

re = BubbleSort(l2)
List.Print(re)