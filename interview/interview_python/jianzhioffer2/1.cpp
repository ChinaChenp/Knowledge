// string 类简单实现

#include <iostream>
#include <utility>
#include <algorithm>
#include <string.h>

using std::cout;
using std::endl;
using std::swap;

class String{
    public:
        // 默认构造函数
        String():m_str(new char[1]){
            m_str[0] = '\0';
        }

        String(const char *str):m_str(new char[strlen(str) + 1]) {
            strcpy(m_str, str);
        }

        // 拷贝构造函数
        String(const String & oth):m_str(new char[oth.size() + 1]) {
            strcpy(m_str, oth.c_str());
        }

        // 析构函数
        ~String() {
            delete []m_str;
        }

        // 赋值拷贝函数
        String & operator=(String th) {
            swap(th);
            return *this;
        }

        // 交换
        void swap(String & th) {
            std::swap(m_str, th.m_str);
        }

        // 长度
        int size() const {
            return strlen(m_str);
        }

        // 原始字符串
        const char * c_str() const {
            return m_str;
        }
    private:
        char *m_str;
};



int main(int argc, char *argv[]) {
    String a;
    std::cout << a.c_str() << std::endl;

    String b("chenp");
    std::cout << b.c_str() << ", "<< b.size() << std::endl;

    String c(b);
    std::cout << c.c_str() << ", " << c.size() << std::endl;
    return 0;
}