// resize and resverse

#include <vector>
#include <iostream>

using namespace std;

class Info {
public:
    Info() : m(0) {
        //std::cout << "默认构造 " << m << "\n";
    };

    Info(int i) : m(i) {
        std::cout << "构造 " << m << "\n";
    };

    ~Info() {
        std::cout << "析构 " << m << "\n";
    };

    int Get() const {
        return m;
    };
private:
    int m;
};

void print(std::vector<Info> &ve) {
    for (auto & s : ve) {
        std::cout << s.Get() << ", ";
    }
    std::cout << "\n";
}

void resize() {
    std::vector<Info> ve;
    for (int i = 0; i < 10; i++) {
        ve.emplace_back(Info(i));
    }
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    print(ve);
    // size=10, cap=16

    ve.resize(8);
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    print(ve);
    // size=8, cap=16 会触发析构

    ve.resize(20, Info(20));
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    print(ve);
    // size=20, cap=32 默认填充新值
}

void reserve() {
    std::vector<Info> ve;
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    ve.reserve(6); // 第7个元素会产生重复拷贝
    for (int i = 0; i < 10; i++) {
        ve.emplace_back(Info(i));
        std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    }
    print(ve);
    std::cout << "~~~~size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    // size=10, cap=12

    ve.reserve(8);
    print(ve);
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    // size=10, cap=12 变小什么都不做

    ve.reserve(20);
    print(ve);
    std::cout << "size=" << ve.size() << ", cap=" << ve.capacity() << std::endl;
    // size=10, cap=20 变大会触发析构，重新拷贝数据到新的块
}

int main(int argc, char *argv[]) {
    //resize();
    reserve();
    return 0;
}