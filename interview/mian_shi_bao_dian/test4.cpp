#include <iostream>
using namespace std;

class A {
private:
    int m_i;
    int m_j;
public:
    A(int a):m_j(a), m_i(m_j){}
    int get_i(){return m_i;}
    int get_j(){return m_j;}
};

int main() {
    A a(98);
    cout << "i=" << a.get_i() << ", j=" << a.get_j() << endl;
}