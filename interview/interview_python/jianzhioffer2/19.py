# 删除链表的节点

# 思路：把下一个节点拷贝到当前节点，把下一个节点删除
# 注意下一个节点为空和是头结点的问题

class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next
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

def Print(head):
    while head != None:
        print('{}->'.format(head.value), end='')
        head = head.next
    print()     

a, b, c, d = Node(1), Node(2), Node(3), Node(4)
head = a
head.next = b
head.next.next = c
head.next.next.next = d
Print(head)

print('del 3')
DeleteNode(head, c)
Print(head)
print('del 4')
DeleteNode(head, d)
Print(head)
print('del 1')
DeleteNode(head, a)
Print(head)
print('del 2')
DeleteNode(head, b)
Print(head)



        
