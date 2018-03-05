cars = ['audi', 'bmw', 'subaru', 'toyota']

#if 条件使用
for car in cars: 
    if car == "bmw":
        print(car.upper())
    else:
        print(car)

#数字比较
age = 18
if 20 > age:
    print("more than age", age)
else:
    print("null")

#使用多个条件 and
if age >= 17 and age <=20:
    print("17<=age<=20")
else:
    print("null")

if age >=18 or age <=7:
    print("age>=18 or age<=7")
else:
    print("null")

#检查特定字段是否包含
print("subaru" in cars)
print("chenp" not in cars)

age = 20
if age < 4:
    print("age < 4")
elif age < 40:
    print("age < 40")
else:
    print("age null")

#strs = ['audi', 'bmw', 'subaru', 'toyota']
strs = []
for str in strs:
    print(str)