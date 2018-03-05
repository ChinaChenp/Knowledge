
alien = {'color': 'green', 'points': 5}
print(alien["color"])

#添加
alien["chenping"] = "100w"
print(alien)

#修改
alien["chenping"] = "150w"
print(alien["chenping"])

#删除
del alien["chenping"]
print(alien)

#遍历
for k, v in alien.items():
    print("key=", k, ",value=", v)

#遍历主键
favorite_languages = { 'jen': 'python', 'sarah': 'c', 'edward': 'ruby', 'phil': 'python'}
for name in favorite_languages.keys():
    print(name)

for language in favorite_languages.values():
    print(language)

#排序
sorts = {"1":"11", "5":"55", "3":"33", "6":"66", "2":"22", "4":"44"}
for key in sorted(sorts.keys()):
    print(key)

for value in sorted(sorts.values()):
    print(value)