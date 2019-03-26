# 链表中环的入口节点

class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next

# 判断链表是否有环，慢指针走一步，快指针走两步
def MeetingNode(head):
    if head == None:
        return None

    slow = head.next
    if slow == None:
        return None

    # 快指针每次走2步，慢指针每次走1步
    fast = slow.next 
    while slow != None and fast != None:
        if slow == fast:
            return slow

        slow = slow.next 
        fast = fast.next 
        if fast != None:
            fast = fast.next 
    return None

# 计算环中节点数量
def CountNode(head):
    if head == None:
        return 0

    cur = head
    count = 1
    # 要从1开始计算，并且用cur.next,因为cur等于head
    while cur.next != head:
        count += 1
        cur = cur.next 
    return count

# 计算环入口节点
# 思路：判断是否有环->计算环中节点个数k->快指针先走k步
def EntryNodeOfLoop(head):
    #先判断是否有环 没有换不需要判断
    meeting_node = MeetingNode(head)
    if meeting_node == None:
        return None
    print(meeting_node.value)
    #计算环中节点个数
    num = CountNode(meeting_node)
    print(num)
    #快慢指针
    fast = slow = head
    while num > 0:
        fast = fast.next
        num -= 1
    
    while fast != slow:
        fast = fast.next 
        slow = slow.next 
    return slow 


a, b, c, d, e = Node(1), Node(2), Node(3), Node(4), Node(5)
head = a
head.next = b
head.next.next = c
head.next.next.next = d
head.next.next.next.next = e
head.next.next.next.next.next = c # 3这个节点会出现环

re = EntryNodeOfLoop(head)
print(re.value)
