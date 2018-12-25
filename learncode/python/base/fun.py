#简单函数
def greet_user():
    print("hello!")

greet_user()

#传参
def greet_user1(username):
    print("hello,", username.title(), "!")
greet_user1("chenp")

#位置实参
def describe_pet(animal_type, pet_name='dog'): 
    """显示宠物的信息"""
    print("\nI have a " + animal_type + ".")
    print("My " + animal_type + "'s name is " + pet_name.title() + ".")
describe_pet('hamster', 'harry')
describe_pet('dog', 'willie')

#关键字实参
describe_pet(pet_name='harry', animal_type='hamster')

#默认实参
describe_pet('shit')

#以下调用方式均可
def describe_pet1(pet_name, animal_type='dog'): 
    """显示宠物的信息"""
    print("\nI have a " + animal_type + ".")
    print("My " + animal_type + "'s name is " + pet_name.title() + ".")
describe_pet1('willie') 
describe_pet1(pet_name='willie')
describe_pet1('harry', 'hamster') 
describe_pet1(pet_name='harry', animal_type='hamster') 
describe_pet1(animal_type='hamster', pet_name='harry')

#返回值
def get_formatted_name(first_name, last_name):
    full_name = first_name + ' ' + last_name
    return full_name.title()
name = get_formatted_name("chen",'ping')
print(name)

#多值返回
def get_multi_return():
    a = 5
    b ="shit"
    return a, b
a, b = get_multi_return()
print(a, b)

#传递列表
def greet_users(users):
    for user in users:
        msg = "hello," + user.title() + "!"
        print(msg)
users = ['hannah', 'ty', 'margot']
greet_users(users)

#修改列表,users默认传递引用
def modify_users(users):
#def modify_users(users[:]): # 禁止修改
    for i in range(len(users)):
        users[i] = users[i].title()
        print(users[i])
users1 = ['hannah', 'ty', 'margot']
modify_users(users1)
print(users1)

#传递任意参数
def make_pizza(*toppings):
    print(toppings)
    print(toppings[0])

make_pizza("a", "b", "~!@")
make_pizza("5", 1)

#任意数量的关键字实参
def build_userinfo(first, last, **userinfo):
    infos = {}
    infos["first_name"] = first
    infos["last_name"] = last

    for k, v in userinfo.items():
        infos[k] = v
    return infos
re = build_userinfo("chenp", "ping", xunbo='laohu', wangyi='yilan')
print(re)
