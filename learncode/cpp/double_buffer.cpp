#include <stdint.h>
#include <stdio.h>
#include <memory>
#include <thread>
#include <atomic>
#include <iostream>
#include <chrono>
 
class Config {
public:
    Config(int &v):val(v){};
    ~Config(){};
    int Get() {return val;}
private:
    int val;
};

int inc_val = 0;

class DataCenter {
public:
    DataCenter() {
    }

    // 程序启动的时候由主线程加载完成
    void init() {
        index = 0;
        switchover[index].reset(new Config(inc_val));
    }
    // *) 切换函数, 由更新线程调用
    void swith() {
        auto cur_index = (index.load() == 0) ? 1 : 0;
        auto old_index = index.load();
        printf("----------cur=%d, old=%d-----------\n", cur_index, old_index);
     
        // *) 新资源ready   
        inc_val++;
        switchover[cur_index].reset(new Config(inc_val));
        // *) 正式切换
        index = cur_index;
         
        // *) 旧资源reset, 引入计数减一
        std::chrono::milliseconds dura(500);
        std::this_thread::sleep_for(dura);
        switchover[old_index].reset();
    }
 
    // *) 访问资源, 由前端线程调用
    std::shared_ptr<Config> getConfig() {
        return switchover[index.load()];
    }
 
private:
    std::atomic<int> m_index;
    // *) 切换数组
    std::shared_ptr<Config> m_switchover[2];
};

void getVal(std::shared_ptr<DataCenter> ptr) {
    int last = -1;
    while (true) {
        auto v = ptr->getConfig();
        auto val = v->Get();
        if (val != last) {
            printf("now=%d, old=%d\n", val, last);
            last = val;
        }
    }
}

void setVal(std::shared_ptr<DataCenter> ptr) {
    while (true) {
        // std::chrono::milliseconds dura(5000);
        // std::this_thread::sleep_for(dura);
        ptr->swith();
        ptr->swith();
    }
}

int main() {
    std::shared_ptr<DataCenter> data(new DataCenter);
    data->init();

    std::thread threads[10];
    for (auto & t: threads) {
        t = std::thread(getVal, data);
    }

    auto v = std::thread(setVal, data);

    for (auto & t : threads) {
        t.join();
    }

    v.join();

    return 0;
}