#include <iostream>
#include "func.h"

int main() {
    std::cout << "Hello, World!" << std::endl;
    int c = add(1, 2);
    std::cout << "add 1 + 2 = " << c << std::endl;
    return 0;
}