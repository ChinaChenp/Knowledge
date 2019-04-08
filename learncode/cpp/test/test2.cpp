// cpp 返回一个对象会触发那些操作

#include <iostream>
using namespace std;

struct cpp_obj{
    cpp_obj() {
        std::cout << "ctor" << std::endl;
    }

    cpp_obj(const cpp_obj & c) {
        std::cout << "copy ctor" << std::endl;
    }

    cpp_obj & operator=(const cpp_obj & c) {
        std::cout << "operator=" << std::endl;
        return *this;
    }

    ~cpp_obj() {
        std::cout << "dtor" << std::endl;
    }
};

cpp_obj return_obj() {
    cpp_obj b;
    std::cout << "before return" << std::endl;
    return b;
}

int main() {
    cpp_obj a;
    a = return_obj();
    return 0;
}