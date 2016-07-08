#linux 命令行使用
## 目录索引

- [基础](#基础)
- [日常使用](#日常使用)
- [vim命令](#vim命令)
- [文本处理](#文本处理)
- [系统调试](#系统调试)
- [单行脚本](#单行脚本)


## 基础  
- 查看文件尾部20行 `tail filename -n 20`  

- 查看文件头尾20行 `head filename -n 20`  

- 在根目录下查找chenp文件 `find  / -name  "chenp" `

- 创建一个3.c的软连接指向2.c `ln -s 2.c 3.c`  

- 统计某个目录下文件个数 `ls pathname | wc -l`  

- 查看磁盘
  - 查看磁盘情况 `df -h` 或者 `df -h ./`  
  - 查看目录每个文件大小 `du -sh` 或者 `du * -sh`  

- 查看当前路劲挂载那个分区 `df -a path`

- 搜索一个目录下所有的关键字 `grep -rn 'chenp' ./*.cc`

- 查找并替换一个目录下的所有字符串 ``` sed -i "s/oldstr/newstr/g" `grep old -rl /www` ```

- date命令
  - 显示当前时间戳  `date +%s`
  - 显示某一刻时间戳
  ```
  date -d '20140929 10:11:11' +%s
  date -d '2014-09-29 10:11:11' +%s
  date -d 20140929 +%s   //29号零点时间戳
  ```
  - 把日期转换成时间戳  `date -d @1411956671`


- scp服务器之间文件拷贝,夸网段无法使用
  - 从服务器上下载文件夹
  ```bash
  scp  -r username@servername:/path/filename
  ```
  - 上传本地文件到服务器
  ```bash
  scp -r /path/filename username@servername:/path
  ```

- nc服务器之间文件拷贝，跨网段也可以使用，比如说有跳板机
  - 在源服务器上启动文件监控服务(目标服务器ip:10.168.99.62)
  `nc -l 1234 < test.txt`
  - 在目标服务器上获取文件
  `nc 10.168.99.62 1234 > test.txt`

## 日常使用
- xz文件解压
```bash
1、xz -d   *.tar.xz  
2、tar -vxf  *.tar
```
- 

## vim命令    
- 打开文件并跳到指定行 `vim +3 main.c 或者 vim main.c +3`  


## 文本处理

## 系统调试  
- 导出线某个进程堆栈信息 `gstack pid > gstack.log` 或者 `pstack pid > pstack.log`

- 查看cpu信息
CPU的核心数是指物理上，也就是硬件上存在着几颗物理cpu,指的是真实存在是cpu处理器的个数，1个代表一颗2个代表2颗cpu处理器。
*核心数*：一个核心就是一个物理线程,英特尔有个超线程技术可以把一个物理线程模拟出两个线程来用,充分发挥CPU性能，意思是一个核心可以有多个线程。
*线程数*：线程数是一种逻辑的概念，简单地说，就是模拟出的CPU核心数。比如，可以通过一个CPU核心数模拟出2线程的CPU，也就是说，这个单核心的CPU被模拟成了一个类似双核心CPU的功能。
  1.查看物理cpu个数
  `grep 'physical id' /proc/cpuinfo | sort -u | wc -l`
  2.查看核心数量
  `grep 'core id' /proc/cpuinfo | sort -u | wc -l`
  3.查看线程数
  `grep 'processor' /proc/cpuinfo | sort -u | wc -l`

- 查看某个进程启动时间  `ps -eo pid,lstart  | grep  pid`  

- `cat query_result.txt   |  sort -k2,2 -k1,1n `  先按第二列排序，在按照第一列数字排序

## 单行脚本

