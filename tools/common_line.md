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


- 服务器之间文件拷贝
  - 从服务器上下载文件夹
  ```bash
  scp  -r username@servername:/path/filename
  ```
  - 上传本地文件到服务器
  ```bash
  scp -r /path/filename username@servername:/path
  ```



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

## 单行脚本

