//  下面输出什么

#include <iostream>

using namespace std;

struct A {
    int b;
    char c;
    long d;
    short e;
};

int func() {
    char a = 120;
    char b = a + 8;
    printf("%d\n", a + 8);
    printf("%d", b);
    return 0;
}

int main() {
    int arr[] = {11, 12, 13, 14, 15};
    int *ptr = arr;

    *(ptr++) += 100;
    printf("%d, %d\n", *ptr, *(++ptr));

    for (int i = 0; i < 5; i++) {
        printf("%d-, ", arr[i]);
    }
    printf("\n");

    printf("sizeof A = %lu\n", sizeof(A));
    printf("sizeof int = %lu\n", sizeof(int));
    printf("sizeof char = %lu\n", sizeof(char));
    printf("sizeof long = %lu\n", sizeof(long));
    printf("sizeof short = %lu\n", sizeof(short));

    func();

    return 0;
}