def fib(n):
    out = []
    a, b = 0, 1
    while a < n:
        out.append(a)
        a, b = b, a + b
    return out

re = fib(50)
print(re)