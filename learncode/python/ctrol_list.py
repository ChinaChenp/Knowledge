#遍历列表
slice = ['alice', 'david', 'carolina']
for s in slice:
    print(s)

#range函数，左包含
for i in range(1, 7):
    print(i)

#使用range创建list
numbers = list(range(1, 9))
print(numbers)

#指定步长打印
even_numbers = list(range(2, 15, 2))
print(even_numbers)

#求连续平方数
squares = []
for i in range(0, 11):
    squares.append(i* i)
print(squares)

#求最大、最小值
print(max(numbers), min(numbers), sum(numbers))

#列表解析方法
squares1 = [v**2 for v in range(0, 11)]
print(squares1)

#切片操作
players = ['charles', 'martina', 'michael', 'florence', 'eli']
print(players[:3])
print(players[1:4])
print(players[2:])
print(players[-2:])

#遍历切片
for player in players[0:3]:
    print(player.title())

#复制
#cars2 = players #引用
cars2 = players[:] #复制
print("cars2=", cars2)

players[0]="charless"
print(cars2)
print(players)