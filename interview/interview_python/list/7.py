# 判断链表是否是循环链表
# 思路：快慢指针，快指针一次走2步，慢指针一次走一步
import list_common as List

def IsLoop(head):
    if head == None:
        return False

    slow = head
    fast = slow.next
    if fast == None:
        return False
    
    while fast != None and slow != None:
        if fast == slow:
            print(fast.value, slow.value)
            return True
        fast = fast.next 
        slow = slow.next 

        if fast != None:
            fast = fast.next 
    return False

a, b, c, d, e, f, g, h = List.Node(1), List.Node(2), List.Node(3), List.Node(4), List.Node(5), List.Node(6), List.Node(7), List.Node(8)
l2 = a
l2.next = c
l2.next.next = e
l2.next.next.next = g
l2.next.next.next.next = h

# l1 = b
# l1.next = d
# l1.next.next = f
# l1.next.next.next = g
# l1.next.next.next.next = h
# l1.next.next.next.next.next = a
# l1.next.next.next.next.next.next = f

# 要分开测试，python是公用一个地址
#print(IsLoop(l1))
print(IsLoop(l2))
