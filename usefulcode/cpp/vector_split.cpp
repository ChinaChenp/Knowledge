
std::vector<std::vector<int64_t>> splitVec(const std::vector<int64_t> &all, uint32_t max_size) {
    uint32_t n = all.size() / max_size;
    uint32_t vec_n = all.size() % max_size == 0 ? n : n + 1;
    std::vector<std::vector<int64_t>> sub_vec(vec_n);

    unsigned int total_size = 0;
    for (uint32_t i = 0; i < n; i++) {
        for (uint32_t j = i*max_size; j < (i+1)*max_size; j++) {
            sub_vec[i].push_back(all[j]);
        }
        total_size += sub_vec[i].size();
    }

    if (total_size < all.size()) {
        for (uint32_t j = total_size; j < all.size(); j++) {
            sub_vec[n].push_back(all[j]);
        }
    }
    return sub_vec;
}