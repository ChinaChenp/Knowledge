# 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
# 备注：重复的数字是挨着的 比如：1->2->3->3->4->4->5 输出：1->2->5
# 思路一：开辟空间缓存一个set
# 思路二：顺序遍历N2复杂度

import list_common as List

def delSameKey(head):
    if head == None:
        return None
    
    cur = head
    while cur != None:
        pre = cur
        next_node = cur.next 

        # 相同节点都是连续的,从当前节点往后找相同节点把相同节点删掉
        while next_node != None: 
            if cur.value == next_node.value:
                pre.next = next_node.next #删除相同节点
            else:
                pre = next_node  #调整节点
            next_node = next_node.next
        cur = cur.next 
    return head

l2 = List.Node(1)
l2.next = List.Node(2)
l2.next.next = List.Node(3)
l2.next.next.next = List.Node(3)
l2.next.next.next.next = List.Node(4)
l2.next.next.next.next.next = List.Node(4)
l2.next.next.next.next.next.next = List.Node(9)
l2.next.next.next.next.next.next.next = List.Node(5)
l2.next.next.next.next.next.next.next.next = List.Node(5)

re = delSameKey(l2)
List.Print(re)