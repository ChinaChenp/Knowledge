#广度优先搜索

from collections import deque

#判断最后位是否为‘m’ 标识苹果商
def person_is_seller(name):
    #print(name, name[-1])
    return name[-1] == 'm'

def graph_search(mp):
    search_queue = deque()
    search_queue += mp["you"]

    searched = {}
    while search_queue:
        person = search_queue.popleft()
        if person not in searched:
            if person_is_seller(person):
                print('{} is a mango seller'.format(person))
                return True
            else:
                search_queue += mp[person]
                searched[person] = ""
    return False


graph = {}
graph["you"] = ["alice", "bob", "claire"]
graph["bob"] = ["anuj", "peggy"]
graph["alice"] = ["peggy"]
graph["claire"] = ["thom", "jonny"]
graph["anuj"] = []
graph["peggy"] = []
graph["thom"] = []
graph["jonny"] = []

print(graph_search(graph))