#链表反转

import list_common as List

def list_reverse(head):
    if head == None or head.next == None:
        return head
    
    q, p = head, head.next 
    head.next = None
    while p:
        tmp = p.next
        p.next = q
        q = p
        p = tmp
    head = q  # 最后p已经指向none了
    return head

head = List.Node(1)
head.next = List.Node(2)
head.next.next = List.Node(3)
head.next.next.next = List.Node(4)

List.Print(head)
new_list = list_reverse(head)
List.Print(new_list)