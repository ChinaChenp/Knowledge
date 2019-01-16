#include <iostream>
#include <cstdint>

using namespace std;

typedef void (*Fun)();   //定义一个函数指针类型变量类型 Fun
typedef int64_t intType; //32位机器使用int32_t

class Base {
  public:
    Base() { cout << "Base construct" << endl; }
    virtual void f() { cout << "Base::f()" << endl; }
    virtual void g() { cout << "Base::g()" << endl; }
    virtual void h() { cout << "Base::h()" << endl; }
    virtual ~Base() {cout << "~Base{}" << endl;}
};

class Derive : public Base {
  public:
    Derive() { cout << "Derive construct" << endl; }
    virtual void f1() { cout << "Derive::f1()" << endl; }
    virtual void g1() { cout << "Derive::g1()" << endl; }
    virtual void h1() { cout << "Derive::h1()" << endl; }
    virtual ~Derive() { cout << "~Derive{}" << endl; }
};

// 打印虚函数表地址，逐个调用虚函数
int test1() {
    Base *b = new Base();
    cout << "首地址：" << *(intType *)(&b) << endl;

    Fun funf = (Fun)(*(intType *)*(intType *)b);
    Fun fung = (Fun)(*((intType *)*(intType *)b + 1)); //地址内的值 即为函数指针的地址，将函数指针的地址存储在了虚函数表中了
    Fun funh = (Fun)(*((intType *)*(intType *)b + 2));

    funf();
    fung();
    funh();

    cout << (Fun)(*((intType *)*(intType *)b + 4)) << endl; //最后一个位置为0 表明虚函数表结束 +4是因为定义了一个 虚析构函数

    delete b;
    return 0;
}

// 一般继承无虚函数覆盖
int test2() {
    Derive *b = new Derive();
    cout << "首地址：" << *(intType *)(&b) << endl;

    Fun funf = (Fun)(*(intType *)*(intType *)b);
    Fun fung = (Fun)(*((intType *)*(intType *)b + 1)); 
    Fun funh = (Fun)(*((intType *)*(intType *)b + 2));

    Fun funf1 = (Fun)(*(intType *)*(intType *)b + 3);
    Fun fung1 = (Fun)(*((intType *)*(intType *)b + 4));
    Fun funh1 = (Fun)(*((intType *)*(intType *)b + 5));

    funf();
    fung();
    funh();

    cout << "drive" << endl;
    funf1();
    fung1();
    funh1();

    //cout << (Fun)(*((intType *)*(intType *)b + 4)) << endl; 

    delete b;
    return 0;
}


int main()
{
    test1();
    test2();
    return 0;
}