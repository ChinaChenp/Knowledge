//
// Created by chenp on 2019-05-06.
//

#include <iostream>
#include <vector>

using namespace std;

std::vector<std::vector<int64_t>> split_vector(std::vector<int64_t> &all, uint32_t max_size) {
    int n = all.size() / max_size;
    int vec_n = all.size() % max_size == 0 ? n : n + 1;
    std::vector<std::vector<int64_t>> sub_vec(vec_n);

    int total_size = 0;
    for (int i = 0; i < n; i++) {
        for (int j = i*max_size; j < (i+1)*max_size; j++) {
            sub_vec[i].push_back(all[j]);
        }
        total_size += sub_vec[i].size();
    }

    if (total_size < all.size()) {
        for (int j = total_size; j < all.size(); j++) {
            sub_vec[n].push_back(all[j]);
        }
    }
    return sub_vec;
}

void Print(std::vector<std::vector<int64_t>> & ve) {
    std::cout << "size=" << ve.size() << std::endl;
    for (auto &i : ve) {
        for (auto &j : i) {
            std::cout << j << ", ";
        }
        std::cout << "\n";
    }
}

int main(int argc, char *argv[]) {
    std::vector<int64_t> ve ={1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16};
    //std::vector<int64_t> ve ={1, 2, 3, 4, 5};
    auto re = split_vector(ve, 5);
    Print(re);
}