#message = input("Tell me something, and I will repeat it back to you: ")
#print(message)

#age = input("how old are you?")
#if int(age) > 20:
#    print("you age is", int(age))

count = 0
while count < 10:
    print(count)
    count += 1

prompt = "\nTell me something, and I will repeat it back to you:" 
prompt += "\nEnter 'quit' to end the program. "

active = True
while active:
    msg = input(prompt)
    if msg == 'quit':
        active = False
    else:
        print(msg)

#删除列表中重复的元素
pets = ['dog', 'cat', 'dog', 'goldfish', 'cat', 'rabbit', 'cat']
print(pets)

while 'cat' in pets:
    pets.remove('cat')
print(pets)