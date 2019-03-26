# 数的N次方

def Power(base, exp):

    # 0的n次方没有意义，0的-n求倒数会出错
    if base >= -1e-9 and base <= 1e-9 and exp < 0:
        return 0.0
    
    # 次方为负数要用正数n次方求倒数
    abs_exp = abs(exp)

    # 计算次方
    re = 1.0
    for _ in range(0, abs_exp):
        re *= base
    
    # 次方为负数要求倒数
    if exp < 0:
        return 1.0 / re
    return re

print(Power(2, 3))
print(Power(1, 2))
print(Power(2, -3))
print(Power(-2, -3))
print(Power(0, -3))
