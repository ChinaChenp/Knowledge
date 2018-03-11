#全部读取
with open("./pi_digits.txt") as file_object:
    contents = file_object.read()
    print(contents)

#逐行读取
with open('pi_digits.txt') as file_object:
    for line in file_object:
        print(line.rstrip())

#使用文件内容的值
with open('pi_digits.txt','r') as file_object:
    lines = file_object.readlines()
pi_string = ''
for line in lines:
    pi_string += line.rstrip()
print(pi_string)

#读取文件
with open('programming.txt', 'w') as file_object:
    file_object.write('I love programming.\n')
    file_object.write("I love creating new games.\n") #写入多行

#追加文件
with open('programming.txt', 'a') as file_object:
    file_object.write("I also love finding meaning in large datasets.\n")
    file_object.write("I love creating apps that can run in a browser.\n")