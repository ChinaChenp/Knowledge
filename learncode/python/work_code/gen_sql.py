
strs = '''alter table passport_openid_{} drop index uniq_idx_openid,add unique key `uniq_idx_openid_source` (`openid`, `source`);'''

for i in range(97):
    #print strs % (i)
    print(strs.format(i))
