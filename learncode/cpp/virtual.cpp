#include <iostream>  
  
using namespace std;  
  
class Base  
{  
public:  
    Base(){cout<<"Base construct"<<endl;}  
    virtual void f() {cout<<"Base::f()"<<endl;}  
    virtual void g() {cout<<"Base::g()"<<endl;}  
    virtual void h() {cout<<"Base::h()"<<endl;}  
    virtual ~Base(){}  
};  
  
int main()  
{  
    typedef void (*Fun)();  //定义一个函数指针类型变量类型 Fun  
    Base *b = new Base();  
    //虚函数表存储在对象最开始的位置  
    //将对象的首地址输出  
    cout<<"首地址："<<*(int*)(&b)<<endl;  
  
    Fun funf = (Fun)(*(int*)*(int*)b);  
    Fun fung = (Fun)(*((int*)*(int*)b+1));//地址内的值 即为函数指针的地址，将函数指针的地址存储在了虚函数表中了  
    Fun funh = (Fun)(*((int *)*(int *)b+2));  
  
    funf();  
    fung();  
    funh();  
  
    cout<<(Fun)(*((int*)*(int*)b+4))<<endl; //最后一个位置为0 表明虚函数表结束 +4是因为定义了一个 虚析构函数  
  
    delete b;  
    return 0;  
} 