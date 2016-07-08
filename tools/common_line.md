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

