# 两个栈模拟一个队列

# 大概思路：stack1作为入队，stack2作为出队
# 1)入队的时候直接压入stack1
# 2)出队的时候判断stack2是否为空，不为空直接返回栈顶，否则把stack1组个压入stack2返回栈顶

class Queue:
    def __init__(self, max_size):
        self.stack1 = []
        self.max_size1 = max_size
        self.stack2 = []
        self.max_size2 = max_size
    def push(self, v):
        if self.max_size1 == len(self.stack1):
            return False
        self.stack1.append(v)
    def pop(self):
        if len(self.stack2) != 0:
            return self.stack2.pop()
        if len(self.stack1) == 0:
            return None

        while len(self.stack1):
            v = self.stack1.pop()
            self.stack2.append(v)
        return self.stack2.pop()

q = Queue(5)
q.push(1)
print('push 1')
q.push(2)
print('push 2')
print('pop', q.pop())
q.push(3)
print('push 3')
q.push(4)
print('push 4')
print('pop', q.pop())
q.push(5)
print('push 5')
print('pop', q.pop())
print('pop', q.pop())
q.push(6)
print('pop', q.pop())
print('pop', q.pop())
