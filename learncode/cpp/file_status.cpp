#include <iostream>
#include <stdio.h>
#include <unistd.h>
#include <fstream>

bool checkFileValid(const std::string &file_name) {
    int ret = access(file_name.c_str(), F_OK | R_OK);
    if (ret != 0) {
        return false;
    }

    std::ifstream file(file_name.c_str());
    if (file.peek() == std::ifstream::traits_type::eof()) {
        return false;
    }
    return true;
}

int main(int argc, char *argv[]) {
    auto ret = checkFileValid(argv[1]);
    ::printf("out %d", ret);
}