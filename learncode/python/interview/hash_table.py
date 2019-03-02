#散列表

class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value

class Map:
    def __init__(self, hash_size, hash=hash):
        self.m_solt = [[] for _ in range(hash_size)]
        self.m_hash_size = hash_size
        self.hash = hash
    #插入
    def add(self, key, value):
        node = Node(key, value)
        hash_key = self.hash(key) % self.m_hash_size
        for v in self.m_solt[hash_key]:
            if v.key == key:
                v.value = value
                return
        self.m_solt[hash_key].append(node)

    #读取
    def get(self, key):
        hash_key = self.hash(key) % self.m_hash_size
        for v in self.m_solt[hash_key]:
            if v.key == key:
                return v.value
        return None
    
    #删除
    def remove(self, key):
        hash_key = self.hash(key) % self.m_hash_size

        index = 0
        for v in self.m_solt[hash_key]:
            if v.key == key:
                break
            index = index + 1
        del self.m_solt[hash_key][index]
    #打印
    def print(self):
        print("-----------------------------")
        for s in self.m_solt:
            for v in s:
                print('key={}, v={}'.format(v.key, v.value))

mp = Map(100)
names = ["chenping", "xunbo", "ceshi", "liubing", "golang"]
for v in range(len(names)):
    mp.add(names[v], v)
mp.print()

for v in names:
    re = mp.get(v)
    print(re)
mp.print()

mp.get("fdsafds")
mp.print()

mp.remove("xunbo")
mp.print()

mp.add("chenping", 11)
mp.print()




