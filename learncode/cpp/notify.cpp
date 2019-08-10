#include <unistd.h>
#include <sys/epoll.h>
#include <sys/inotify.h>
#include <string>
#include <map>
#include <thread>

using namespace std;

typedef int (* FILE_WATCHER_FUNC)(std::string & file_name, void* arg);

struct FileWatcher{
    std::string file_name;
    FILE_WATCHER_FUNC watcher_func;
    void* arg;
};

class FileNotify{
public:
    FileNotify():is_init_(false){}
    int Init();
    int Trigger();

    void AddFileWatcher(const std::string & file_name,
            FILE_WATCHER_FUNC watcher_func, void* arg);

    static FileNotify* GetInstance(){
        return s_fs_notify_instance;
    }

private:
    bool is_init_;
    int epoll_fd_;
    static FileNotify* s_fs_notify_instance;
    std::map<int, FileWatcher*> fd_info_map_;
};


FileNotify* FileNotify::s_fs_notify_instance = new FileNotify();

#define EPOLL_MAX_NUM 256
#define MAX_EVENT_NUM 128
#define RELOAD_BUFFER_MAX_LEN 1000

int FileNotify::Init(){
    if (!__sync_bool_compare_and_swap(&is_init_, false, true)) {
        return 0;
    }

    epoll_fd_ = epoll_create(EPOLL_MAX_NUM);
    return 0;
}

int FileNotify::Trigger(){
    struct epoll_event events[MAX_EVENT_NUM];
    char buffer[RELOAD_BUFFER_MAX_LEN];

    int nfds = epoll_wait(epoll_fd_, events, MAX_EVENT_NUM, 0);

    for(int i=0; i<nfds; ++i){
        int fd = events[i].data.fd;
        //探测是否有更新
        int ret = read(fd, buffer, RELOAD_BUFFER_MAX_LEN);
        if(-1 == ret) {
            continue;
        }

        std::map<int, FileWatcher*>::iterator idx = fd_info_map_.find(fd);
        if(fd_info_map_.end() == idx) {
            continue;
        }

        FileWatcher* file_watcher = idx->second;
        inotify_add_watch(fd, file_watcher->file_name.c_str(), IN_MODIFY);
        file_watcher->watcher_func(file_watcher->file_name, file_watcher->arg);
        //file_watcher->watcher_func(file_watcher->file_name);
    }
    return 0;
}

//加入需要监测的文件
//此处应在框架Init层进行，故不做并发保护
void FileNotify::AddFileWatcher(const std::string & file_name,
        FILE_WATCHER_FUNC watcher_func, void* arg){
    Init();
    int fd = inotify_init();
    inotify_add_watch(fd, file_name.c_str(), IN_MODIFY);
    struct epoll_event ev;
    ev.events = EPOLLIN | EPOLLET;
    ev.data.fd = fd;
    epoll_ctl(epoll_fd_, EPOLL_CTL_ADD, fd, &ev);

    FileWatcher* file_watcher = new FileWatcher();
    file_watcher->file_name = file_name;
    file_watcher->watcher_func = watcher_func;
    file_watcher->arg = arg;
    fd_info_map_[fd] = file_watcher;
    return;
}

int print(std::string & file_name, void* arg) {
    printf("ok %s\n", file_name.c_str());
}

void CheckFile() {
    while(true) {
  	FileNotify::GetInstance()->Trigger();
        sleep(1);
    }
}

int main() {
    std::thread t(CheckFile);
    FileNotify::GetInstance()->AddFileWatcher("./1.txt", print, (void*)NULL);
    FileNotify::GetInstance()->AddFileWatcher("./2.txt", print, (void*)NULL);
    FileNotify::GetInstance()->AddFileWatcher("./3.txt", print, (void*)NULL);

    sleep(1000);
    printf("out\n");
}