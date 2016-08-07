#《c++ interview汇总》
## 第一部分：基础知识  

## 第二部分：网络知识  
- tcp/ip分层
  - ![Alt text](/picture/tcp_ip.jpg "tcp/ip分层")
  - 链路层
    - 也称作数据链路层或网络接口层（在第一个图中为网络接口层和硬件层），通常包括操作系统中的设备驱动程序和计算机中对应的网络接口卡。它们一起处理与电缆（或其他任何传输媒介）的物理接口细节。
    - ARP（地址解析协议）和RARP（逆地址解析协议）是某些网络接口（如以太网和令牌环网）使用的特殊协议，用来转换IP层和网络接口层使用的地址。

  - 网络层
      - 也称作互联网层（在第一个图中为网际层），处理分组在网络中的活动，例如分组的选路。在TCP/IP协议族中，网络层协议包括IP协议（网际协议），ICMP协议（Internet互联网控制报文协议），以及IGMP协议（Internet组管理协议）。
      - IP是一种网络层协议，提供的是一种不可靠的服务，它只是尽可能快地把分组从源结点送到目的结点，但是并不提供任何可靠性保证。同时被TCP和UDP使用。TCP和UDP的每组数据都通过端系统和每个中间路由器中的IP层在互联网中进行传输。
      - ICMP是IP协议的附属协议。IP层用它来与其他主机或路由器交换错误报文和其他重要信息。
      - IGMP是Internet组管理协议。它用来把一个UDP数据报多播到多个主机。

  - 传输层
      - 主要为两台主机上的应用程序提供端到端的通信。在TCP/IP协议族中，有两个互不相同的传输协议：TCP（传输控制协议）和UDP（用户数据报协议）。
      - TCP为两台主机提供高可靠性的数据通信。它所做的工作包括把应用程序交给它的数据分成合适的小块交给下面的网络层，确认接收到的分组，设置发送最后确认分组的超时时钟等。由于运输层提供了高可靠性的端到端的通信，因此应用层可以忽略所有这些细节。为了提供可靠的服务，TCP采用了超时重传、发送和接收端到端的确认分组等机制。
      - UDP则为应用层提供一种非常简单的服务。它只是把称作数据报的分组从一台主机发送到另一台主机，但并不保证该数据报能到达另一端。一个数据报是指从发送方传输到接收方的一个信息单元（例如，发送方指定的一定字节数的信息）。UDP协议任何必需的可靠性必须由应用层来提供。

  - 应用层
      - 应用层负责处理特定的应用程序细节。

## 第三部分：操作系统  

## 第四部分：算法实现  
- [在一个数组中找出重复数值](/interview/src/001.c)
- [字符串反转](/interview/src/002.c)
- [翻转句子中单词](/interview/src/003.c)
- [输入一个单向链表,输出该链表中倒数第k个结点](/interview/src/004.c)
- [判断字符串是否是对称的即(回文字符串)](/interview/src/005.c)
- [求子数组的最大和](/interview/src/006.c)
- [在一个字符串中找到第一个只出现一次的字符](/interview/src/007.c)
- [输入一个表示整数的字符串,把该字符串转换成整数并输出](/interview/src/008.c)
- [设计包含min 函数的栈](/interview/src/009.c)
- [字符串的左旋转操作](/interview/src/010.c)
- [字符串是否包含问题](/interview/src/011.c)
- [linux C 所有字符串操作函数实现](/interview/src/012.c)
- [求和1+2+3+...N](/interview/src/013.c)
- [输入n,用最快的方法求Fibonacci数列的第n项](/interview/src/014.c)
- [求整数二进制表示中1的个数](/interview/src/015.c)
- [输入整数n,计算从1到n这n个整数的十进制表示中1出现的次数和](/interview/src/016.c)
- [如何对n个数进行排序,要求时间复杂度O(n),空间复杂度O(1)](/interview/src/017.c)
- [(*)和为n连续正数序列](/interview/src/018.c)
- [10分钟写出二分法查找,并调试通过](/interview/src/019.c)
- [(*)数组中有一个数字出现的次数超过了数组长度的一半,找出这个数字](/interview/src/020.c)
- [从尾到头输出链表/倒序输出单向链表](/interview/src/021.c)
- [(*)在字符串中删除特定的字符](/interview/src/022.c)
- [(*)单链表就地逆置,头插法](/interview/src/023.c)
- [(*)在字符串中找出连续最长的数字串,并把这个串的长度返回](/interview/src/024.c)
- [求一个数组的最长递减子序列,和(024)类似](/interview/src/025.c)
- [调整数组顺序使奇数位于偶数前面](/interview/src/026.c)
- [在O(1)时间内删除链表结点](/interview/src/027.c)
- [已知一个字符串,比如asderwsde,寻找其中的一个子字符串比如sde 的个数,如果没有返回0,有的话返回子字符串的个数](/interview/src/028.c)
- [字符串原地压缩,题目描述:"eeeeeaaaff" 压缩为"e5a3f2",请编程实现](/interview/src/029.c)
- [有一个由大小写组成的字符串，现在需要对他进行修改，将其中的所有小写字母排在大写字母的前面](/interview/src/030.c)
- [函数将字符串中的字符'*'移到串的前部分，前面的非'*'字符后移，但不能改变非'*'字符的先后顺序,函数返回串中字符'*'的数量。](/interview/src/031.c)
- [(*)求两个串中的第一个最长子串(神州数码以前试题).如"abractyeyt","dgdsaeactyey"的最大子串为"actyey".](/interview/src/032.c)
- [(*)两个有序单链表合并为一个有序的单链表](/interview/src/033.c)
- [编写一个函数求一个数组中的第二大数](/interview/src/034.c)
- [时钟的时分秒针每天重叠几次？时分每天相遇几次？分秒每小时遇到几次？](/interview/src/035.c)
- [输出字符串的所有排列](/interview/src/036.cpp)
- [数值的整数次方.实现函数double Power(double base, int exponent),求base 的exponent次方,不需要考虑溢出](/interview/src/037.c)
- [找出两个数之间的素数的个数,并打印](/interview/src/038.c)
- [ip合法性检查](/interview/src/039.c)
- [最大公约数问题(编程之美2.7)](/interview/src/040.c)
- [寻找数组中的最大值和最小值(编程之美2.10)](/interview/src/041.c)
- [字符串移位包含的问题(s1是否可以由s2移位后得到)](/interview/src/042.c)
- [各种排序算法实现](/interview/src/043.c)
- [圆圈中最后剩下的数字](/interview/src/044.cpp)
- [寻找第1500个丑数(因子只包含2,3,5的数)](/interview/src/045.c)
- [洗牌算法](/interview/src/046.cpp)
- [数组中全是(0, 1),写一个算法求最长连续1的起始位置和终止位置](/interview/src/047.cpp)
- [输入n个整数，找出其中最小的k个数](/interview/src/048.cpp)
- [随机生成和为S的N个正整数](/interview/src/049.cpp)
- [去除重复字符并排序](/interview/src/050.cpp)
- [寻找最大（最小）的K个数(topk问题)](/interview/src/051.cpp)
- [leetcode 2.1.1 Remove Duplicates from Sorted Array](/interview/src/052.c)
- [leetcode 2.1.2 Remove Duplicates from Sorted Array II ](/interview/src/053.c)
- [leetcode 2.1.3 Sear in Rotated Sorted Array](/interview/src/054.c)
