#include <string.h>
#include <stdio.h>
#include <iostream>

using namespace std;

class A {
    public:
    A(){m_a=1;}
    ~A(){}

    void func(){printf("A %d\n", m_a);}

    int m_a;
};

class B: public A
{
  public:
    B() { m_a = 2;}
    ~B() {}

    //void func() {printf("B %d\n", m_a); }

    int m_a;
};

int main() {
    B b;
    b.func();
    printf("a=%d\n", b.m_a);
}