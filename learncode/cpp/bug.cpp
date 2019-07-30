
void copy(char *src) {
    int len = strlen(src);
    char *dst = new char(len + 1);
    memset(dst, 0, len + 1);
    strcpy(dst, src);
    delete []dst;
}

int main() {
    char a[10] = {'a', 'b', 'c', '\0'};
    copy(a);
}