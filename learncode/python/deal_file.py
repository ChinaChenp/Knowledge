#使用文件内容的值
with open('data2.txt','r') as file_object:
    lines = file_object.readlines()

count_limit = 0
all_str = []
for line in lines:
    mp = {}
    strs = line.split("||")
    name = strs[0].split(" ")[1]
    mp["name"] = name

    tmp = ""
    query_words = ""
    activityId = ""
    all_ok = 0
    for s in strs:
        if s.startswith("req_common"):
            tmp = s
            all_ok = all_ok + 1
        if s.startswith("activityId"):
            activityId = s.split("=")[1]
            all_ok = all_ok + 1
        if s.startswith("req_query"):
            tm = s.split("=")
            if len(tm) != 2:
                continue
            query_words = tm[1]
            all_ok = all_ok + 1
        if all_ok == 3:
            break
    vecs = tmp.split(",")
    for v in vecs:
        kv = v.split(":")
        if len(kv) != 2:
            continue
        k = kv[0].strip()
        v = kv[1].strip()
        mp[k] = v
    count_limit = count_limit + 1
    if count_limit == 5000:
        break
    #print('name=%s, city=%s, lng=%s, lat=%s' % (mp["name"], mp["city"], mp["lng"], mp["lat"]))

    #if mp["name"] == "RecByWordsReq":
    #    s = '{"cityID":%s, "lat":%s, "lng":%s, "pageNo":%s, "count":%s, "queryWords":"%s"}' % (mp["city"], mp["lat"], mp["lng"], mp["pageno"], mp["count"], query_words)
    #    all_str.append(s)
    if mp["name"] == "BannerRecShopReq":
        s = '{"cityID":%s, "lat":%s, "lng":%s, "pageNo":%s, "count":%s, "queryWords":"%s", "ActivityID":"%s"}' % (mp["city"],
         mp["lat"], mp["lng"], mp["pageno"], mp["count"], query_words, activityId)
        all_str.append(s)
doc_str = ",\n"
re = doc_str.join(all_str)
print('[%s]' % re)