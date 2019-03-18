# 从尾打印链表
 
class Node:
    def __init__(self, value, next = None):
        self.value = value
        self.next = next

def ReverseList(head):
    if head == None or head.next == None:
        return

    p = head
    q = head.next

    head.next = None
    while q:
        tmp = q.next
        q.next = p
        p = q
        q = tmp
    head = p
    return head

head = Node(1)
head.next = Node(2)
head.next.next = Node(3)
head.next.next.next = Node(4)
head.next.next.next.next = Node(5)

re = ReverseList(head)
print(re.value, re.next.value, re.next.next.value,
      re.next.next.next.value, re.next.next.next.next.value)

