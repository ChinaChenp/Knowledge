// 不同命名空间转换

#include <iostream>
#include <algorithm>
#include <vector>

//using namespace std;
namespace A {
    typedef struct names {
        std::string name;
        int age;
    }Names;

    typedef struct txt {
        std::vector<int> ve;
    }TXT;

    typedef struct info {
        int a;
        double b;
        Names *names;
        TXT txt;
    }Info;
};

namespace B
{
    typedef struct names {
        std::string name;
        int age;
    }Names;

    typedef struct txt{
        std::vector<int> ve;
    } TXT;

    typedef struct info {
        int a;
        double b;
        Names *names;
        TXT txt;
    }Info;
}; // namespace A

int main(int argc, char *argv[]) {
    A::Info *info = new(A::Info);
    info->a = 1;
    info->b = 1.2;
    info->names = new(A::Names);
    info->names->age = 30;
    info->names->name = "chenping";
    std::vector<int> ve = {1, 2, 3, 4};
    info->txt.ve = ve;

    B::Info *info1 = new(B::Info);
    info1->names = new(B::Names);

    //std::swap(info, info1);
    //std::swap(*info, *info1);
    info1 = (B::Info*)info;
    std::cout << "a " << info1->a << std::endl;
    std::cout << "b " << info1->b << std::endl;
    std::cout << "names->age " << info1->names->age << std::endl;
    std::cout << "names->name " << info1->names->name << std::endl;
    for (auto & c: info1->txt.ve) {
        std::cout << c << ", ";
    }
    std::cout << std::endl;

    return 0;
}