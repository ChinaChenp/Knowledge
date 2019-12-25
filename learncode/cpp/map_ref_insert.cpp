#include <iostream>
#include <unordered_map>
#include <map>
#include <string>

using namespace std;

int main() {
    std::map<std::string, int32_t> mp;
    mp["1"] = 1;
    mp["2"] = 2;

    auto & v = mp["1"];
    v += 10;
    mp["2"] += 10;

    for (auto & m : mp) {
        std::cout << m.first << "=" << m.second << ", ";
    }
    std::cout << std::endl;
}