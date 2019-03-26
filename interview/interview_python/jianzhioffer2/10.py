#菲波那切数列

#递归
def Fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return Fibonacci(n - 1) + Fibonacci(n - 2)


#非递归实现
def Fibonacci1(n):
    a, b = 0, 1

    for _ in range(0, n):
        a, b = b, a + b
    return a

print(Fibonacci(15), Fibonacci1(15))