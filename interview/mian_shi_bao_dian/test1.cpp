#include <string.h>
#include <stdio.h>
#include <iostream>

using namespace std;

class A {
    public:
    A(){m_a=1, m_b=2;}
    ~A(){}

    void func(){printf("A %d,%d", m_a, m_b);}

    int m_a;
    int m_b;
};

class B
{
  public:
    B() { m_c = 3;}
    ~B() {}

    void func() {printf("B %d", m_c); }

    int m_c;
};

int main() {
    A a;
    B *pb = (B*)&a;
    pb->func();
}