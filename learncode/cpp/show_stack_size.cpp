// 测试栈大小

#include <stdio.h>
#include <stdlib.h>

void show_stack() {
    struct rlimit lmt;
    getrlimit(RLIMIT_STACK,&lmt);
    printf("rlim_cur=%u,rlim_max=%u\n",lmt.rlim_cur,lmt.rlim_max);
}

void test(int i) {
    char temp[1024*1024] = {0}; // 1M
    
    temp[0] = i;
    temp[0] ++;
    printf("%s %d num = %d!\r\n", __FUNCTION__, __LINE__, temp[0]);
    test(temp[0]);
}

int main(void) {
    show_stack();
    test(0);
    return 0;
}