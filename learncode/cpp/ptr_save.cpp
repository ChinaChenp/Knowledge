#include <iostream>

int main() {
    // std::cout << "Hello, World!" << std::endl;
    // //int c = add(1, 2);
    // std::cout << "add 1 + 2 = " << c << std::endl;

    int64_t *ptr = NULL;
    int64_t **save = &ptr;
    int64_t a = 5;
    std::cout << "a addr = " << &a << std::endl;
    std::cout << "ptr addr = " << &ptr << std::endl;
    std::cout << "ptr dsc addr = " << ptr << std::endl;

    *save = &a;
    std::cout << save << std::endl;
    std::cout << *save << std::endl;
    std::cout << **save << std::endl;

    std::string str = "chenping";
    str.append("yes");
    std::cout << str << std::endl;
    return 0;
}