//
// Created by chenp on 2019-05-06.
//

#include <iostream>

#include <string>

#include <chrono>

#include <thread>

#include <future>

using namespace std::chrono;

std::string fromDb(std::string recvData) {
    std::this_thread::sleep_for(seconds(5));
    return "DB_" + recvData;
}

std::string fromFile(std::string recvData) {
    std::this_thread::sleep_for(seconds(5));
    return "File_" + recvData;
}

int main() {
    auto start = system_clock::now();
    std::cout << "start read db data" << std::endl;
    std::future < std::string > db = std::async (std::launch::async, fromDb, "DbData");
    std::cout << "end read db data" << std::endl;

    std::cout << "start read file data" << std::endl;
    std::future < std::string > file = std::async (std::launch::async, fromFile, "FileData");
    std::cout << "end read file data" << std::endl;

    db.wait();
    file.wait();

    auto end = system_clock::now();
    auto diff = duration_cast < std::chrono::seconds > (end - start).count();
    std::cout << "Total Time taken= " << diff << "Seconds" << std::endl;

    auto db_data = db.get();
    auto file_data = file.get();

    std::string data = db_data + " :: " + file_data;

    std::cout << "Data = " << data << std::endl;

    return 0;
}