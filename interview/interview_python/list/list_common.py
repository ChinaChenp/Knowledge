#list基础组件

def Print(head):
    cur = head
    while cur != None:
        print('{}->'.format(cur.value), end='')
        cur = cur.next
    print()

class Node:
    def __init__(self, value, next=None):
        self.value = value
        self.next = next