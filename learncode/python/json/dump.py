import json

filename = 'numbers.json'

numbers = ['1', '2', '3', '4', '5']
with open(filename, 'w') as file_object:
    json.dump(numbers, file_object)


with open(filename) as file_object:
    numbers2 = json.load(file_object)
    print(numbers2)
    for number in numbers2:
        print(number)