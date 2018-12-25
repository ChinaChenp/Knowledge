a = set('abracadabra')
b = set("alacazam")

print('a=', a)
print('b=', b)

print('a-b=', a-b)
print('a|b=', a|b)
print('a&b=', a&b)
print('a^b=', a^b)

if 'a' in a:
    print("a in a")

if 'x' not in a:
    print("x not in a")