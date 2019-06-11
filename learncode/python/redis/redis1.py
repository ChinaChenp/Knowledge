# redis 简单使用

from redis import StrictRedis
import time

redis = StrictRedis(host='10.179.16.222', port=6379, db=0, password='')
re = redis.get('chenp')
print(re)

for i in range(0, 100):
    key = 'chenp' + str(i)
    redis.set(key, i, 86400)

time.sleep(0.5)

for i in range(0, 100):
    key = 'chenp' + str(i)
    re = redis.get(key)
    if re == None:
        print(key, 'is none')
    else:
        print(re)