# 单向链表中倒数第K个节点

# 思路：三种异常情况
# 1) head 为空
# 2）k大于链表长度
# 3）k == 0


class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next

def FindKthToTail(head, n):
    if head == None or n == 0:
        return None
    
    # 快指针先走N步
    fast = slow = head
    for _ in range(0, n):
        if fast != None:
            fast = fast.next
        else:
            return None # 异常情况链表不够长
    
    # 快慢指针同时走
    while fast:
        fast = fast.next
        slow = slow.next
    return slow.value


a, b, c, d = Node(1), Node(2), Node(3), Node(4)
head = a
head.next = b
head.next.next = c
head.next.next.next = d
print(FindKthToTail(head, 1))
print(FindKthToTail(head, 3))
