bicycles = ["trek", 'cannondale', 'redline', 'specialized'] 
print(bicycles)

#循环遍历列表
for l in bicycles:
    print(l)

#指定打印某个列表值
print(bicycles[0])

print(bicycles[-1])
print(bicycles[-2])

#修改列表
bicycles[0] = 'chenping'
print(bicycles)

#追加列表
bicycles.append("xunbo")
print(bicycles)

#插入列表
bicycles.insert(1, "chenping2")
print(bicycles)

#删除列表
del bicycles[2]

#弹出列表
last = bicycles.pop()
print(bicycles)
print(last)

bicycles.pop(0)
print(bicycles)

#删除第一个出现的字符串
bicycles.remove("redline")
print(bicycles)

#永久排序
cars = ['bmw', 'audi', 'toyota', 'subaru']
#cars.sort()
cars.sort(reverse=True)
print(cars)

#临时排序
cars1 = ['bmw', 'audi', 'toyota', 'subaru']
print(sorted(cars1))

#翻转
cars1.reverse()
print(cars1)

#长度
re = len(cars1)
print(re)

# motorcycles = ['honda', 'yamaha', 'suzuki'] 
# print(motorcycles[3])

#相加
l1 = [1, 2, 3, 4, 5, "chenp"]
l2 = [11, 34, "ping"]
l3 = l1 + l2
print(l3)

#乘法
l1 = [45] * 10
print(l1)

#判断
if "chenp" in l1:
    print("chenp in ok")
else:
    print("chenp not in")

#插入
numbers = [1, 5]
print(numbers)
numbers[0:0] = [1, 2, 3, 4]
print(numbers)

#列表推导式
vec = [1, 2, 3]
re = [3*x for x in vec]
print(re)

ll = list([1, 2, 3, 4, 5])
ll.insert(0, 11)
print(ll)
ll.pop(0)
print(ll)