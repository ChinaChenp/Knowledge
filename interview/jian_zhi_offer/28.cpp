#include <iostream>
using namespace std;

void Permutaion(char *str, char *begin) {
    if (*begin == '\0') {
        std::cout << str << std::endl;
    } else {
        for (char *ch = begin; *ch != '\0'; ch++) {
            std::swap(*begin, *ch);
            Permutaion(str, begin+1);
            std::swap(*begin, *ch);
        }
    }
}

int main() {
    char str[] = "abc";
    Permutaion(str, str);
}
