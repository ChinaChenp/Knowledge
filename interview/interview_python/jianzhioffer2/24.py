#链表反转

class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next


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

head = Node(1)
head.next = Node(2)
head.next.next = Node(3)
head.next.next.next = Node(4)

new_list = list_reverse(head)
print(new_list)
print(new_list.value, new_list.next.value,
      new_list.next.next.value, new_list.next.next.next.value)
