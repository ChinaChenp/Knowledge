# 目录
> 主要记录推荐、搜索、广告方向下架构设计、核心原理、试用场景等

## 第一部分：推荐 & 搜索 & 广告
  - 第一章：名词解释
    - 1.1 [推荐名词](docs/section1/1.1.md)
    - 1.2 [搜索名词](docs/section1/1.2.md)
    - 1.3 [广告名词](docs/section1/1.3.md)
  - 第二章：基础知识
    - 2.1 [加解密](docs/section1/2.1.md)
    - 2.2 url encode & decode
    - 2.3 yaml & toml 
  - 第三章：整体架构
    - 3.1 搜索
    - 3.2 推荐
    - 3.3 广告
    - 3.4 账号
  - 第四章：模块设计
    - 4.1 broker
    - 4.2 es
    - 4.3 recall
    - 4.4 rank
    - 4.5 rerank
    - 4.5 fc
    - 4.6 lbs
      - geohash
      - mmap
      - epoll 文件fd
      - 射线法
    - 4.7 ip定位
    
## 第二部分：基础架构 & 组件
  - nginx
  - thrift
    - 相比其他http、ip直连优势
  - redis
    - [高可用搭建集群](https://mp.weixin.qq.com/s/Z-PyNgiqYrm0ZYg0r6MVeQ)
    - 如何容错、宕机
    - 如何迁移
  - ice
  - odin
  - baimai
  - hive
  - MapReduce
  - mq 
    - 高可用
    - 消费积压怎么做
  - mysql
    - 高可用集群方案
    - binlog
  - leveldb
  - lru & lfu
  - zookeeper
  - rpc
  - 服务发现
  - 负载均衡
  
## 第三部分：工程化 & 稳定性
  - 第一章：稳定性
    - 1.1 [全局意识](http://naotu.baidu.com/file/bb3cefe050e5e894c2dfaa699267aae7)
    - 1.2 熔断
    - 1.3 降级
    - 1.4 限流
    - 1.5 监控
    - 1.6 报警
    - 1.7 大盘
    - 1.8 异地多活
    - 1.9 [复盘](docs/section3/1.9.md)
  - 第二章：工程化
    - 运维
    - 上线
    
## 第四部分：流程 & 能效 & 团队建设
  - 1.1 上线流程
  - 1.2 值班
  - 1.3 工具化
  - 1.4 测试流程
  - 1.5 研发过程
  - 1.6 [周会](docs/section4/1.6.md)
  - 1.7 周报
  - 1.8 团队建设
  - 1.9 绩效考核
  - 1.10 面试

## 第五部分：工具心得
  - git
  - mysql
  - vim
  - [shell](docs/section5/1.4.md)
  
## 第六部分：基本功
  - 第一章：语言
    - 1.1 [Go](docs/section6/1.1.md)
    - 1.2 [C++](docs/section6/1.2.md)
  - 第二章：算法
  - 第三章：操作系统
  - 第四章：网络
  - 第五章：方案设计
    - 5.1 [id生成器](docs/section6/5.1.md)
  
## 第七部分：经验 & 入坑
  - 2018.4.35

## 第八部分：阅读 & 思考
  - 1.1 [2018-10](docs/1.1.md)
  > 思考

```
 _                  _                      
| |__  _   _    ___| |__   ___ _ __  _ __  
| '_ \| | | |  / __| '_ \ / _ \ '_ \| '_ \ 
| |_) | |_| | | (__| | | |  __/ | | | |_) |
|_.__/ \__, |  \___|_| |_|\___|_| |_| .__/ 
       |___/                        |_|    
```
