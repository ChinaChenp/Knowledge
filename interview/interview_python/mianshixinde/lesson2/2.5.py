#跳台阶问题

# 首先考虑最简单的情况。如果只有1级台阶，那显然只有一种跳法。如果有2级台阶，那就有两种跳的方法了：一种是分两次跳，每次跳1级；另外一种就是一次跳2级。
# 现在我们再来讨论一般情况。我们把n级台阶时的跳法看成是n的函数，记为f(n)。

# 当n>2时，第一次跳的时候就有两种不同的选择：
# 一是第一次只跳1级，此时跳法数目等于后面剩下的n-1级台阶的跳法数目，即为f(n-1)；
# 另外一种选择是第一次跳2级，此时跳法数目等于后面剩下的n-2级台阶的跳法数目，即为f(n-2)。
# 因此n级台阶时的不同跳法的总数f(n)=f(n-1)+f(n-2)。

def Fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    if n == 2:
        return 2
    return Fibonacci(n - 1) + Fibonacci(n - 2)

def Fibonacci1(n):
    a, b = 0, 1
    for _ in range(0, n):
        a, b = b, a + b
    return b


print(Fibonacci(21))
print(Fibonacci1(21))