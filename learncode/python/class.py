class Dog():
    def __init__(self, name, age):
        self.name = name
        self.age = age
    def sit(self):
        print(self.name.title(), "is sitting now.")
    def roll_over(self):
        print(self.name.title(), "is roll over.")

dog = Dog("dahua", 12)
dog.sit()
dog.roll_over()
print(dog.name, dog.age)


class Car():
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year
        self.odometer_reading = 0 #默认参数
    def get_discriptive_name(self):
        lname = str(self.year) + ' ' + self.model + " " + self.make
        return lname
    def get_odometer_reading(self):
        return self.odometer_reading
    def update_odometer(self, mileage):
        if mileage >= self.odometer_reading:
            self.odometer_reading = mileage
        else:
            print("You can't roll back an odometer!")
    def increment_odometer(self, mileage):
        self.odometer_reading += mileage
    def fill_gas_tank(self):
        print(self.make.title(), "is full ok")
car = Car("audi", "a4", 2016)
print(car.get_discriptive_name())
print(car.get_odometer_reading())

#方法1-修改参数
car.odometer_reading = 15
print(car.get_odometer_reading())

#方法2-通过方法修改
car.update_odometer(16)
print(car.get_odometer_reading())

#增量更新
car.increment_odometer(100)
print(car.get_odometer_reading())

#类继承
class ElectricCar(Car):
    def __init__(self, make, model, year):
        super().__init__(make, model, year) #初始化父类,继承父类所有方法
        self.battery_size = 123 #定义电动车特有的属性
    def describe_battery(self):
        print("this car has a " + str(self.battery_size) + "-kwh battery.")
    def fill_gas_tank(self): # 重写方法
        print("This car doesn't need a gas tank!")
my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_discriptive_name())
print(my_tesla.describe_battery())
print(my_tesla.fill_gas_tank())