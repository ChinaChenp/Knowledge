#include <iostream>
#include <string>

using namespace std;
struct Info {
    Info() {
        std::cout << "copy" << std::endl;
    }
    int a;
    std::string b;
};

void test() {
    const int a = 10;  ////a是常量 C中的常量是假的，C++中是真的
    int *p      = (int *)&a;
    *p          = 20;  //改变的是临时开辟的temp变量
    std::cout << "a = " << a << std::endl;
    std::cout << "*p = " << *p << std::endl;

    Info infos;
    infos.a = 100;
    infos.b = "chenping";

    const auto &c = infos;

    std::cout << "c.a=" << c.a << ", c.b=" << c.b << "addr=" << &infos << std::endl;
    Info *infos_p = (Info *)&c;
    infos_p->a    = 200;
    infos_p->b    = "niubi";
    std::cout << "c.a=" << c.a << ", c.b=" << c.b << "addr=" << infos_p << std::endl;
}

int main() {
    test();

    return 0;
}