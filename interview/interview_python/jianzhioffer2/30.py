# 最小mini栈

# 思路：用一个辅助栈每次压入时跟栈顶比较
# 比栈顶小：直接压入
# 比栈顶大: 把栈顶重新压入

class StackMin:
    def __init__(self):
        self.stack1 = []
        self.stack2 = [] # 辅助栈
    def push(self, v):
        if len(self.stack2) == 0:
            self.stack2.append(v)
        else:
            top = self.stack2[-1]
            if top > v:
                self.stack2.append(v)
            else:
                self.stack2.append(top)
        self.stack1.append(v)

    def pop(self):
        s = self.stack1.pop()
        m = self.stack2.pop()
        return s, m
    def min(self):
        return self.stack2[-1]

s = StackMin()
s.push(3)
print('min:', s.min())
s.push(4)
print('min:', s.min())
s.push(2)
print('min:', s.min())
s.push(1)
print('min:', s.min())
