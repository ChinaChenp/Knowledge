message = "Hello Python Crash Course reader!" 
print('I told my friend, "Python is my favorite language!"')
print("\nhello,world")

print (message.title())
print(message.upper())
print(message.lower())

print("a" + "+" + "b")

print(" shit".lstrip())
print("shit  ".rstrip())
print("  shit  ".strip())

msg = "One of Python's strengths is its diverse community."
print(msg)

for msg in "shit shit fuck shit".split(" "):
    print(msg)

msg1 = "fuck off shit"
for i in enumerate(msg1):
    print(i)

for j in range(len(msg1)):
    print(j)

