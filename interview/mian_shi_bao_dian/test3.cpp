
struct Test {
    Test(int a){}
    Test(){}
    void fun(){}
};

int main() {
    Test a(1);
    a.fun();

    Test b();
    b.fun();
}