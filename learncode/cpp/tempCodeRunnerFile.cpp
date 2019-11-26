//
// Created by chenp on 2019-04-26.
//

#include <functional>
#include <iostream>
#include <memory>
#include <random>
struct Foo {
    Foo()
        : data(0)
    {
    }
    void print_sum(int& n1, int& n2)
    {
        n1 += 1;
        n2 += 2;
        data += n1 + n2;
        std::cout << "n1=" << n1 << ", n2=" << n2 << ", data=" << data << '\n';
    }
    int data;
};

int main(int argc, char* argv[])
{
    using namespace std::placeholders; // 对于 _1, _2, _3...
    // 绑定指向成员函数指针
    std::vector<std::function<void(int&, int&)>> funs;
    Foo foo;
    auto f1 = std::bind(&Foo::print_sum, &foo, _1, _2);
    auto f2 = std::bind(&Foo::print_sum, &foo, _1, _2);
    funs.push_back(f1);
    funs.push_back(f2);

    int a = 1;
    int b = 1;
    for (auto f : funs) {
        f(std::ref(a), b);
    }
}