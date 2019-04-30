// 单例

class Singleton {
public:
    static Singleton *GetInstance() {
        return m_instance;
    }
private:
    static Singleton * m_instance;
};

Singleton *Singleton::m_instance = new Singleton;

int main(int argc, char *argv[]) {
    auto in = Singleton::GetInstance();
    return 0;
}