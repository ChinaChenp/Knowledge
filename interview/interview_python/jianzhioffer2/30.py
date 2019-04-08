# 最小mini栈

# 思路：用一个辅助栈每次压入时跟栈顶比较
# 比栈顶小：直接压入当前值
# 比栈顶大: 把栈顶取出重新压入

class StackMin:
    def __init__(self):
        self.stack1 = []
        self.stack2 = [] # 辅助栈
    def push(self, v):
        # 处理压入最小值
        if len(self.stack2) == 0: #第一次压入无需判断
            self.stack2.append(v)
        else:
            top = self.stack2[-1]
            if top > v:
                self.stack2.append(v)
            else:
                self.stack2.append(top)
        #压入栈值
        self.stack1.append(v)

    def pop(self):
        if len(self.stack1) == 0:
            return None
        s = self.stack1.pop()
        m = self.stack2.pop()
        return s, m
    def min(self):
        if len(self.stack2) == 0:
            return None
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
