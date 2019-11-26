// 迭代器失效

#include <deque>
#include <vector>
#include <list>
#include <iostream>
#include <map>
#include <unordered_map>

using namespace std;

template<typename T>
void test_list(T li) {
    for (auto iter = li.begin(); iter != li.end();) {
        if (*iter == 4) {
            iter = li.erase(iter);
        } else {
            iter++;
        }
    }

    for (auto & s : li) {
        std::cout << s << ", ";
    }
    std::cout << std::endl;
}

template<typename T>
void test_map(T mp) {
    for (auto iter = mp.begin(); iter != mp.end(); ) {
        if (iter->first == 4) {
            mp.erase(iter++);
            //iter = mp.erase(iter); // c++ 11 support
        } else {
            iter++;
        }
    } 
    
    for (auto & s : mp) {
        std::cout << "{" << s.first << ":" << s.second << "}, ";
    }
    std::cout << std::endl;
}

int main() {
    // 测试顺序容器
    std::vector<int> ve = {4, 1, 2, 3, 4, 4, 8, 4, 9, 4};
    test_list(ve);

    std::deque<int> qu = {4, 1, 2, 3, 4, 4, 8, 4, 9, 4};
    test_list(qu);

    std::list<int> li = {4, 1, 2, 3, 4, 4, 8, 4, 9, 4};
    test_list(li);

    // 测试关联容器
    std::map<int, string> mp = {{4, "4"}, {1, "1"}, {2, "2"}, {4, "4"}, {4, "4"}, {9, "9"}};
    test_map(mp);

    std::unordered_map<int, string> mp1 = {{4, "4"}, {1, "1"}, {2, "2"}, {4, "4"}, {4, "4"}, {9, "9"}};
    test_map(mp1);
}