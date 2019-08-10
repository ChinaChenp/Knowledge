#include <iostream>
#include <ctime>
 
using namespace std;
 
std::string formartTime() {
    time_t now = time(NULL);
    tm *t = localtime(&now);

    // 2019-08-07 11:03:13
    char buf[32] = {0};
    cout << "-------" << buf << "----------" << endl;
    snprintf(buf, sizeof buf, "%04d-%02d-%02d %02d:%02d:%02d",
        1900 + t->tm_year,
        t->tm_mon,
        t->tm_mday,
        t->tm_hour,
        t->tm_min,
        t->tm_sec
        );
    return buf;
}

int main( )
{
    auto date = formartTime();
    cout << "--------" << date << "-----" << endl;
}