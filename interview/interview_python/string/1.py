#判断一个字符串是否包含另外一个字符串

def SubStr(s1, s2):
    l1, l2 = len(s1), len(s2)
    i, j = 0, 0

    while i < l1 and j < l2:
        if s1[i] == s2[j]:
            i += 1
            j += 1
        else:
            i = i - j + 1
            j = 0

    if l2 == j:
        return True
    return False

re = SubStr("fjdfeijfchenpufda", "chenp")
print(re)

re = SubStr("fjdfeijfchenpufda", "chenp1")
print(re)

re = SubStr("fjdfeijfchenpufda", "pufd")
print(re)