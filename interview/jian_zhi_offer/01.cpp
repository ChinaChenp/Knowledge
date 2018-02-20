#include <iostream>
#include  <utility>
#include  <algorithm>
#include <string.h>

using namespace std;

class String {
public:
    String(){}
    ~String(){}
    String & operator=(const String & th);
    String(const String & th);
private:
    char * str;
};

#if 0
String & String::operator=(const String & th) {
    if (this == &th ){
        return *this;
    }

    delete []str;
    str = NULL;
    str = new char[strlen(th.str) + 1];
    strcpy(str, th.str);
    return *this;
}
#else
String & String::operator=(const String & th) {
    if (this != &th ){
        String tmp = String(th);
        std::swap(tmp.str, str);
    }

    return *this;
}
#endif

int main() {

}
