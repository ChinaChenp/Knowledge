# 删除链表的节点

# 思路：把下一个节点拷贝到当前节点，把下一个节点删除
# 注意下一个节点为空和是头结点的问题

import list_common as List

def DeleteNode(head, toDel):
    if head == None or toDel == None:
        return head

    # 如果不是尾节点
    if toDel.next != None:
        toDel.value = toDel.next.value
        toDel.next = toDel.next.next;

    # 如果只有一个节点
    elif head.value == toDel.value:
        head.next = None
    # 多个节点且在尾部，从头开始删除
    else:
        tmp = cur = head
        while cur.value != toDel.value: # 语言问题拷贝过后指针已经发生变化，所以用值代替
            tmp = cur
            cur = cur.next
            print('------', tmp.value, cur.value, toDel.value)
        tmp.next == None  

a, b, c, d = List.Node(1), List.Node(2), List.Node(3), List.Node(4)
head = a
head.next = b
head.next.next = c
head.next.next.next = d
List.Print(head)

print('del 3')
DeleteNode(head, c)
List.Print(head)
print('del 4')
DeleteNode(head, d)
List.Print(head)
print('del 1')
DeleteNode(head, a)
List.Print(head)
print('del 2')
DeleteNode(head, b)
List.Print(head)



        
